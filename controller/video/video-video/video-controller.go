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
	user2 "timtubeWebAdmin/io/user/user"
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
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			util.CreateSession("message", "The video is Too large for now!", app, r)
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/video/video/", 301)
			return
		}
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
		videoObject := domain.Video{"", title, date, now.Format(util.YYYYMMDDHHMMSS), description, isPrivate, price, ""}
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
		//fileSize := strconv.Itoa(handler.Size)
		videoData := domain.VideoData{videoResult.Id, []byte{}, fileBytes, extension, "10"}

		go sendVideo(videoData)

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
func sendVideo(video domain.VideoData) {
	_, err := user.CreateVideoData(video)
	if err != nil {
		fmt.Println("error creating video")
		fmt.Println(err)
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
		var videoOwner domain.User
		var videoOwnerAccount domain.UserAccount
		var role string
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		videoToview, err := user3.ReadVideo(id)
		if err != nil {
			fmt.Println(err, " error reading video!")
		}
		isProcessing := util.IsProcessing(videoToview.DateUploaded)
		if isProcessing == true {
			fmt.Println("still processing")
			util.CreateSession("message", "Data still processing, try again after 10 minutes.", app, r)
			http.Redirect(w, r, "/video/video", 301)
			return
		}
		videoData, err := user.ReadVideoData(videoToview.Id)
		if err != nil {
			fmt.Println(err, " error reading video data")
		}
		userVideoObject, err := userVideo.ReadUserVideoWithVideoId(videoToview.Id)
		if err != nil {
			fmt.Println(err, " error reading user video data")
		} else {
			videoOwnerAccount, err = user4.ReadUserAccount(userVideoObject.CustomerId)
			if err != nil {
				fmt.Println(err, " error reading user Account data")
			} else {
				videoOwner, err = user2.ReadUser(videoOwnerAccount.Email)
				if err != nil {
					fmt.Println(err, " error reading user data")
				}
				role, err = util.GetRoleName(videoOwner.RoleId)
				if err != nil {
					fmt.Println(err, " error reading ROle data")
				}
			}
		}

		sEnc := base64.StdEncoding.EncodeToString(videoData.Picture)

		message := util.GetSession("message", app, r)
		base := "http://74.208.50.103:8081/"
		type PageData struct {
			Name        string
			Surname     string
			Video       domain.Video
			Message     string
			Picture     string
			Email       string
			Base        string
			VideoData   domain.VideoData
			UserVideo   domain.UserVideo
			UserAccount domain.UserAccount
			User        domain.User
			Role        string
		}
		data := PageData{name, surname, videoToview, message, sEnc, email, base, videoData, userVideoObject, videoOwnerAccount, videoOwner, role}
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
		if userRole != "superAdmin" {
			userId := util.GetUserId(app, r)
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
		} else {
			videos, _ = user3.ReadVideos()
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
