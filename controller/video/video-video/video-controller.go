package video_video

import (
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	"timtubeWebAdmin/io/pcloud"
	user4 "timtubeWebAdmin/io/user/user-account"
	userVideo "timtubeWebAdmin/io/user/user-video"
	user3 "timtubeWebAdmin/io/video/video"
	user "timtubeWebAdmin/io/video/video-data"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", viewVideos(app))
	r.Get("/video/{id}", viewVideo(app))
	r.Post("/update", updateVideoDetails(app))
	r.Post("/create", createVideo(app))

	return r
}
func checkOnAndOff(isPrivate string) bool {
	if isPrivate == "on" {
		return true
	}
	return false
}
func createVideo(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, handler, err := r.FormFile("video")
		title := r.PostFormValue("title")
		date := r.PostFormValue("date")
		description := r.PostFormValue("description")
		isPrivate := checkOnAndOff(r.PostFormValue("isPrivate"))
		price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)

		email := app.Session.GetString(r.Context(), "email")
		if err != nil && email == "" {
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/video/video/", 301)
			return
		}

		now := time.Now()
		videoObject := domain.Video{"", title, date, now.Format(util.YYYMMDD), description, isPrivate, price, ""}
		videoResult, err := user3.CreateVideo(videoObject)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/video/video/", 301)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		fileExtension := strings.Split(handler.Filename, ".")
		extension := fileExtension[1]
		videoData := domain.VideoData{videoResult.Id, fileBytes, []byte{}, extension}
		_, err = user.CreateVideoData(videoData)
		if err != nil {
			fmt.Println(err, " error creating video Data")
		}

		user, err := user4.ReadUserAccountWithEmail(email)
		if err != nil {
			fmt.Println(err, " error reading user")
		}
		userVideoObject := domain.UserVideo{"", user.CustomerId, videoResult.Id, now.Format(util.YYYMMDD)}
		_, err = userVideo.CreateUserVideo(userVideoObject)
		if err != nil {
			fmt.Println(err, " error creating user-video")
		}
		http.Redirect(w, r, "/video/video", 301)
		return
	}
}

func updateVideoDetails(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		birthdate := r.PostFormValue("birthdate")
		fmt.Println("data: ", birthdate)
		//email := r.PostFormValue("email")
		//name := r.PostFormValue("name")
		//surname := r.PostFormValue("surname")
		//birthdate,_ := time.Parse(util.YYYMMDD, r.PostFormValue("birthdate"))
		//roleId := r.PostFormValue("roleId")
		//if name != "" && email != "" && surname != "" && roleId != "" {
		//	userObject := domain.User{email, name, surname, birthdate, roleId}
		//	_, err := user2.UpdateUser(userObject)
		//	if err != nil {
		//		fmt.Println(err, " error update user")
		//		util.CreateSession("message", "couldn't update User details try again later", app, r)
		//		http.Redirect(w, r, "/user/user/user/"+email, 301)
		//		return
		//	}
		//}
	}

}

func fileWriter(file multipart.File, extension string) {
	tempFile, err := ioutil.TempFile("files/test", "upload-*."+extension)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Println(err, "fail to Upload File")
	} else {
		pcloud.UploadFile(tempFile)
	}
	// return that we have successfully uploaded our file!
	fmt.Println("Successfully Uploaded File")

}

//TODO we need to cupture the size of the video and format...
func viewVideo(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		videoToview, err := user3.ReadVideo(id)
		if err != nil {
			fmt.Println(err, " error reading video!")
		}
		byteArray, err := user.ReadVideoPicture(id)
		if err != nil {
			fmt.Println(err, " error reading video picture!")
		}
		sEnc := base64.StdEncoding.EncodeToString(byteArray)

		message := util.GetSession("message", app, r)
		base := "http://74.208.50.103:8081/"
		type PageData struct {
			Name    string
			Surname string
			Video   domain.Video
			Message string
			Picture string
			Email   string
			Base    string
		}
		data := PageData{name, surname, videoToview, message, sEnc, email, base}
		var files = []string{}
		filesAgent := []string{
			app.Path + "video/video.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "video/video.html",
			app.Path + "template/super-admin-navigator.html",
			app.Path + "template/topbar.html",
		}
		if userRole == "superAdmin" {
			files = fileAdmins
		} else if userRole == "agent" {
			files = filesAgent
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func viewVideos(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var videos = []domain.Video{}
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		//user, err := user2.ReadUser(email)
		//fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		userId := app.Session.GetString(r.Context(), "userId")
		if userId != "" {
			userVideos, err := userVideo.ReadAllOfUserVideo(userId)
			if err != nil {
				fmt.Println(err, " error reading video")
			}
			for _, userVideoObject := range userVideos {
				video, err := user3.ReadVideo(userVideoObject.VideoId)
				if err != nil {
					fmt.Println(err, " error reading video")
				}
				videos = append(videos, video)
			}
		}

		type PageData struct {
			Videos  []domain.Video
			Name    string
			Surname string
			Message string
		}
		data := PageData{videos, name, surname, message}
		var files = []string{}
		filesAgent := []string{
			app.Path + "video/videos.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "video/videos.html",
			app.Path + "template/super-admin-navigator.html",
			app.Path + "template/topbar.html",
		}
		if userRole == "superAdmin" {
			files = fileAdmins
		} else if userRole == "agent" {
			files = filesAgent
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
