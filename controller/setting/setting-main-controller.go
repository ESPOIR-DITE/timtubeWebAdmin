package setting

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	sctc "timtubeWebAdmin/controller/setting/setting-channel-type"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	channel_type "timtubeWebAdmin/io/channel/channel-type"
	user2 "timtubeWebAdmin/io/user/user"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/channel-type", sctc.Home(app))
	return mux
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		user, err := user2.ReadUser(email)
		fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		channelTypes, err := channel_type.ReadChannelTypes()
		if err != nil {
			fmt.Println(err, " error reading channel types")
		}
		type PageData struct {
			Channels []domain.ChannelType
			Name     string
			Surname  string
			Message  string
			Email    string
		}
		data := PageData{channelTypes, name, surname, message, email}
		ts, err := template.ParseFiles(util.RolePageRender(app, userRole, "settings/settings.html")...)
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
