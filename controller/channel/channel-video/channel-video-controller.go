package channel_type

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	user4 "timtubeWebAdmin/io/channel/channel"
	dct "timtubeWebAdmin/io/channel/channel-type"
	dcv "timtubeWebAdmin/io/channel/channel-video"
	user2 "timtubeWebAdmin/io/user/user"
	user3 "timtubeWebAdmin/io/user/user-video"
	user "timtubeWebAdmin/io/video/video"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Get("/{id}", getHandler(app))
	r.Get("/channel/{id}", getByChannelIdHandler(app))
	r.Get("/video/{id}", getByVideoIdHandler(app))
	r.Post("/create", createHandler(app))
	r.Get("/delete/{id}", deleteHandler(app))
	r.Post("/update", updateHandler(app))
	return r
}

func getByVideoIdHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", http.StatusTemporaryRedirect)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
			return
		}
		channels, err := dcv.ReadChannelVideo(id)
		if err != nil {
			fmt.Println(err, "error reading channels type")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels domain.ChannelVideos
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-Type.html")...)
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

func getByChannelIdHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", http.StatusTemporaryRedirect)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
			return
		}
		channels, err := dcv.ReadChannelVideoByChannelId(id)
		if err != nil {
			fmt.Println(err, "error reading channels type")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.ChannelVideos
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-Type.html")...)
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

func getHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", http.StatusTemporaryRedirect)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
			return
		}
		channels, err := dct.ReadChannelType(id)
		if err != nil {
			fmt.Println(err, "error reading channels type")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels domain.ChannelType
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-Type.html")...)
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

func deleteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			_, err := dcv.DeleteChannelVideo(id)
			if err != nil {
				fmt.Println(err, "error creating channel type")
			}
			app.Session.Put(r.Context(), "message", "You have Delete Channel type")
			http.Redirect(w, r, "/channel/type", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not Delete, data missing!")
		http.Redirect(w, r, "/channel/type", 301)
		return
	}
}
func createHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//Id          :=r.PostFormValue("channelId")
		channelId := r.PostFormValue("channelId")
		videoId := r.PostFormValue("videoId")
		description := r.PostFormValue("description")

		if channelId != "" && videoId != "" {
			channelVideoObject := domain.ChannelVideos{"", videoId, channelId, description}
			_, err := dcv.CreateChannelVideo(channelVideoObject)
			if err != nil {
				fmt.Println(err, "error creating channel video")
			}
			app.Session.Put(r.Context(), "message", "You have created a new channel video")
			http.Redirect(w, r, "/channel/type", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/channel/type", 301)
		return
	}
}
func updateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		Id := r.PostFormValue("channelVideoId")
		channelId := r.PostFormValue("channelId")
		videoId := r.PostFormValue("videoId")
		description := r.PostFormValue("description")

		if channelId != "" && videoId != "" && Id != "" {
			channelVideoObject := domain.ChannelVideos{Id, videoId, channelId, description}
			_, err := dcv.CreateChannelVideo(channelVideoObject)
			if err != nil {
				fmt.Println(err, "error updating channel video")
			}
			app.Session.Put(r.Context(), "message", "You have created a new channel video")
			http.Redirect(w, r, "/channel/type", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/channel/type", 301)
		return
	}
}
func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		channels := getChannelsVideosOwners()

		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []util.ChannelsVideosOwners
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-videos.html")...)
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

func getChannelsVideosOwners() []util.ChannelsVideosOwners {
	var CVO []util.ChannelsVideosOwners
	var owner domain.User
	channelVidoes, err := dcv.ReadChannelVideos()
	if err != nil {
		fmt.Println(err, "error reading channels")
		return CVO
	}
	for _, channelVideo := range channelVidoes {
		video, err := user.ReadVideo(channelVideo.VideoId)
		if err != nil {
			fmt.Println(err, "error reading video")
		}
		channelObject, err := user4.ReadChannel(channelVideo.ChannelId)
		if err != nil {
			fmt.Println(err, "error reading CHannel")
		}
		ChannelType, err := dct.ReadChannelType(channelObject.ChannelTypeId)
		if err != nil {
			fmt.Println(err, "error reading Channel type")
		}
		userVideo, err := user3.ReadUserVideoWithVideoId(video.Id)
		if err != nil {
			fmt.Println(err, "error reading user video")
		} else {
			owner, err = user2.ReadUser(userVideo.CustomerId)
			if err != nil {
				fmt.Println(err, "error reading")
			}
		}
		CVO = append(CVO, util.ChannelsVideosOwners{channelObject, ChannelType, owner, video})
	}
	return CVO
}
