package channel_type

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	dct "timtubeWebAdmin/io/channel/channel-type"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Get("/{id}", getHandler(app))
	r.Post("/create", createHandler(app))
	r.Get("/delete/{id}", deleteHandler(app))
	r.Post("/update", updateHandler(app))
	return r
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
			_, err := dct.DeleteChannel(id)
			if err != nil {
				fmt.Println(err, "error creating channel type")
			}
			app.Session.Put(r.Context(), "message", "You have Delete Channel type")
			http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not Delete, data missing!")
		http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
		return
	}
}
func createHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//Id          :=r.PostFormValue("channelId")
		Name := r.PostFormValue("name")
		Description := r.PostFormValue("description")

		if Name != "" {
			channelSubscriptionObject := domain.ChannelType{"", Name, Description}
			_, err := dct.CreateChannelType(channelSubscriptionObject)
			if err != nil {
				fmt.Println(err, "error creating channel Type")
			}
			app.Session.Put(r.Context(), "message", "You have created a new channel Type")
			http.Redirect(w, r, "/setting", 301)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/setting", 301)
		return
	}
}
func updateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		Id := r.PostFormValue("channelId")
		Name := r.PostFormValue("channelId")
		Description := r.PostFormValue("channelId")

		if Name != "" && Id != "" {
			channelSubscriptionObject := domain.ChannelType{Id, Name, Description}
			_, err := dct.UpdateChannelType(channelSubscriptionObject)
			if err != nil {
				fmt.Println(err, "error creating channel Type")
			}
			app.Session.Put(r.Context(), "message", "You have created a new channel Type")
			http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/channel/type", http.StatusTemporaryRedirect)
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
		channels, err := dct.ReadChannelTypes()
		if err != nil {
			fmt.Println(err, "error reading channels")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.ChannelType
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-type.html")...)
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
