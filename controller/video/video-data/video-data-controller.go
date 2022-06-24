package video_data

import (
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strings"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	user3 "timtubeWebAdmin/io/video/video"
	user "timtubeWebAdmin/io/video/video-data"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", viewVideoData(app))
	r.Post("/update", updateVideoDetails(app))
	r.Post("/create", createVideo(app))
	r.Post("/video-update", updateVideoData(app))

	return r
}

func updateVideoData(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var videoObject domain.Video
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			util.CreateSession("message", "The video is Too large for now!", app, r)
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/video/video/", 301)
			return
		}
		file, handler, err := r.FormFile("video")
		videoId := r.PostFormValue("videoId")
		email := app.Session.GetString(r.Context(), "email")
		if err != nil && email == "" {
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/video/video/", 301) //todo: should send to the longin
			return
		}
		if videoId != "" {
			videoObject, err = user3.ReadVideo(videoId)
			if err != nil {
				fmt.Println(err, "error reading video")
				util.CreateSession("message", "This video missing", app, r)
				fmt.Println("erro: ", err)
				http.Redirect(w, r, "/video/video/video/"+videoId, 301)
				return
			}
			_, err := user.DeleteVideoData(videoObject.Id)
			if err != nil {
				fmt.Println(err, "error reading video")
				util.CreateSession("message", "Fail to update. Try again later", app, r)
				fmt.Println("erro: ", err)
				http.Redirect(w, r, "/video/video/video/"+videoId, 301)
				return
			}
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		fileExtension := strings.Split(handler.Filename, ".")
		extension := fileExtension[1]
		//fileSize := strconv.Itoa(handler.Size)
		videoData := domain.VideoData{videoObject.Id, []byte{}, fileBytes, extension, "10"}

		go sendVideo(videoData)

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

func createVideo(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("picture")
		videoId := r.PostFormValue("videoId")
		extension := r.PostFormValue("extension")
		size := r.PostFormValue("size")

		email := app.Session.GetString(r.Context(), "email")
		if err != nil && email == "" {
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/video/video/", 301)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err, " error getting bytes!")
			return
		}

		videoData := domain.VideoData{videoId, fileBytes, []byte{}, extension, size}
		_, err = user.ReadVideoData(videoId)
		if err != nil {
			_, err = user.UpdateVideoData(videoData)
			if err != nil {
				fmt.Println(err, " error updating video Data")
			}
		} else {
			_, err = user.CreateVideoPictureData(videoData)
			if err != nil {
				fmt.Println(err, " error creating video Data")
			}
		}

		http.Redirect(w, r, "/video/video", 301)
		return
	}
}

func updateVideoDetails(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("picture")
		videoId := r.PostFormValue("videoId")
		extension := r.PostFormValue("extension")
		size := r.PostFormValue("size")

		email := app.Session.GetString(r.Context(), "email")
		if err != nil && email == "" {
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/video/video/", 301)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		//fileSize := strconv.Itoa(handler.Size)
		videoData := domain.VideoData{videoId, fileBytes, []byte{}, extension, size}
		_, err = user.UpdateVideoData(videoData)
		if err != nil {
			fmt.Println(err, " error creating video Data")
		}

		http.Redirect(w, r, "/video/video", 301)
		return
	}
}

func viewVideoData(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
