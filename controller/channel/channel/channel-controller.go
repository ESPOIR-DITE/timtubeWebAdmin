package channel

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	user3 "timtubeWebAdmin/io/channel/channel"
	user5 "timtubeWebAdmin/io/channel/channel-subscription"
	user "timtubeWebAdmin/io/channel/channel-type"
	user4 "timtubeWebAdmin/io/user/role"
	user2 "timtubeWebAdmin/io/user/user"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", channelHomeHandler(app))
	r.Get("/{id}", getChannelHandler(app))
	r.Get("/user/{id}", getChannelByUserHandler(app))
	r.Get("/region/{id}", getChannelByRegionHandler(app))
	r.Post("/create", createHandler(app))
	r.Get("/delete/{id}", deleteHandler(app))
	r.Post("/update", updateHandler(app))
	return r
}

func getChannelByRegionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/", 301)
			return
		}
		channels, err := user3.ReadChannelByRegion(id)
		if err != nil {
			fmt.Println(err, "error reading channels")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.Channel
		}
		data := PageData{message, name, surname, role, channels}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channels.html")...)
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
func getChannelByUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/", 301)
			return
		}
		channels, err := user3.ReadChannelByUserId(id)
		if err != nil {
			fmt.Println(err, "error reading channels")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.Channel
		}
		data := PageData{message, name, surname, role, channels}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channels.html")...)
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
func getChannelHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/", 301)
			return
		}
		channels, err := user3.ReadChannel(id)
		if err != nil {
			fmt.Println(err, "error reading channels")
			http.Redirect(w, r, "/channel/", 301)
			return
		}
		sEnc := base64.StdEncoding.EncodeToString(channels.Image)
		channelType, err := user.ReadChannelType(channels.ChannelTypeId)
		if err != nil {
			fmt.Println(err, "error reading channel type")
		}
		channelTypes, err := user.ReadChannelTypes()
		if err != nil {
			fmt.Println(err, "error reading channel types")
		}
		owner, err := user2.ReadUser(channels.UserId)
		if err != nil {
			fmt.Println(err, "error reading user(owner of the channel)")
		}
		message := util.GetSession("message", app, r)
		ownerRole, err := user4.ReadRole(owner.RoleId)
		if err != nil {
			fmt.Println(err, "error reading user(owner of the channel) Role")
		}
		subscriptionNumber, err := user5.CountChannelSubscriptionByChannelId(channels.Id)
		if err != nil {
			fmt.Println(err, "error reading Subscription amount")
		}
		type PageData struct {
			Message            string
			Name               string
			Surname            string
			UserRole           string
			Channels           domain.Channel
			Image              string
			ChannelType        domain.ChannelType
			ChannelTypes       []domain.ChannelType
			User               domain.User
			Role               domain.Role
			SubscriptionAmount int64
			Email              string
		}
		data := PageData{message, name, surname,
			role, channels, sEnc, channelType,
			channelTypes, owner, ownerRole, subscriptionNumber, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel.html")...)
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
			_, err := user3.DeleteChannel(id)
			if err != nil {
				fmt.Println(err, "error creating channel")
			}
			app.Session.Put(r.Context(), "message", "You have Delete a new Channel")
			http.Redirect(w, r, "/channel", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not Delete, data missing!")
		http.Redirect(w, r, "/channel", 301)
		return
	}
}
func createHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if err := r.ParseMultipartForm(2 << 10); err != nil {
			util.CreateSession("message", "The video is Too large for now!", app, r)
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/channel/home", 301)
			return
		}
		file, _, err := r.FormFile("photo")
		Name := r.PostFormValue("name")
		ChannelTypeId := r.PostFormValue("chaneTypeId")
		UserId := r.PostFormValue("userId")
		Region := r.PostFormValue("region")
		description := r.PostFormValue("description")

		email := app.Session.GetString(r.Context(), "email")
		if err != nil && email == "" {
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/channel", 301)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		if Name != "" && UserId != "" && len(fileBytes) > 0 {
			channelObject := domain.Channel{"", Name, ChannelTypeId, UserId, Region, util.GetCurrentDate(), fileBytes, "", description}
			_, err := user3.CreateChannel(channelObject)
			if err != nil {
				fmt.Println(err, "error creating channel")
			}
			app.Session.Put(r.Context(), "message", "You have created a new Channel")
			http.Redirect(w, r, "/channel/home", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/channel", 301)
		return
	}
}
func updateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("photo")
		Id := r.PostFormValue("id")
		Name := r.PostFormValue("name")
		ChannelTypeId := r.PostFormValue("chaneTypeId")
		UserId := r.PostFormValue("userId")
		Region := r.PostFormValue("region")
		email := app.Session.GetString(r.Context(), "email")
		description := r.PostFormValue("description")
		if err != nil && email == "" {
			fmt.Println("erro: ", err)
			http.Redirect(w, r, "/channel/", 301)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		if Id != "" && Name != "" && UserId != "" && len(fileBytes) > 0 {
			channelObject := domain.Channel{Id, Name, ChannelTypeId, UserId, Region, util.GetCurrentDate(), fileBytes, "", description}
			_, err := user3.UpdateChannel(channelObject)
			if err != nil {
				fmt.Println(err, "error updating channel")
			}
			app.Session.Put(r.Context(), "message", "You have updated a new Channel")
			http.Redirect(w, r, "/channel", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not update, data missing!")
		http.Redirect(w, r, "/channel", 301)
		return
	}
}
func channelHomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var channels []domain.Channel
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		var err = errors.New("")
		if role == "admin" {
			channels, err = user3.ReadChannels()
			if err != nil {
				fmt.Println(err, "error reading channels")
			}
		} else {
			channels, err = user3.ReadChannelByUserId(email)
			if err != nil {
				fmt.Println(err, "error reading channels")
			}
		}

		channelTypes, err := user.ReadChannelTypes()
		if err != nil {
			fmt.Println(err, "error reading channelTypes")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message      string
			Email        string
			Name         string
			Surname      string
			UserRole     string
			Channels     []domain.Channel
			ChannelTypes []domain.ChannelType
		}
		data := PageData{message, email, name, surname, role, channels, channelTypes}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channels.html")...)
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
