package user

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/user/user"
	user_account "timtubeWebAdmin/controller/user/user-account"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Mount("/account", user_account.Home(app))
	mux.Mount("/user", user.Home(app))
	return mux
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "index.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
