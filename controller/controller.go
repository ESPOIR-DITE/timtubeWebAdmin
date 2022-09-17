package controller

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	cmc "timtubeWebAdmin/controller/channel"
	setting "timtubeWebAdmin/controller/setting"
	"timtubeWebAdmin/controller/user"
	"timtubeWebAdmin/controller/util"
	video2 "timtubeWebAdmin/controller/video"
	"timtubeWebAdmin/domain"
	"timtubeWebAdmin/io"
	user2 "timtubeWebAdmin/io/user/user"
	user3 "timtubeWebAdmin/io/user/user-account"
	user4 "timtubeWebAdmin/io/user/user-detail"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	//mux.Handle("/", homeHandler(env))
	mux.Handle("/", homeHandler(env))
	mux.Handle("/out", outHandler(env))
	mux.Mount("/user", user.Home(env))
	mux.Mount("/video", video2.Home(env))
	mux.Mount("/channel", cmc.Home(env))
	mux.Mount("/setting", setting.Home(env))

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}

func outHandler(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Session.Clear(r.Context())
		http.Redirect(w, r, "/user/account/login", 301)
		return
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := util.GetSession("email", app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		var userRole string
		var isExistString string

		userAccount, err := user3.ReadUserAccountWithEmail(email)
		if userAccount.Status == false {
			util.CreateSession("message", "Please wait for Admin to activate your Account", app, r)
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		user, err := user2.ReadUser(email)
		if user.RoleId != "" {
			userRole, err = util.GetRoleName(user.RoleId)
			if err != nil {
				fmt.Println(err, "\nerror reading user role")
				util.CreateSession("message", "Please try again later", app, r)
				http.Redirect(w, r, "/user/account/login", 301)
				return
			}
			if userRole == "user" {
				util.CreateSession("message", "You are not allowed", app, r)
				http.Redirect(w, r, "/user/account/login", 301)
				return
			}

		} else {
			util.CreateSession("message", "Please wait for Admin confirmation ti assign role!", app, r)
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		if user.Name != "" && user.Surname != "" {
			app.Session.Put(r.Context(), "name", user.Name)
			app.Session.Put(r.Context(), "surname", user.Surname)
			app.Session.Put(r.Context(), "role", userRole)
			app.Session.Put(r.Context(), "userId", userAccount.CustomerId)
		}
		if userRole == "agent" {
			isExist, err := user4.IsExistUserDetailEmail(email)
			if err != nil {
				fmt.Println(err, "error checkin userDetails")
				isExistString = "error"
			} else if isExist == true {
				isExistString = "exist"
			} else if isExist == false {
				isExistString = "not exist"
			}
		}

		DashBoard, err := io.GetDashboardData()
		if err != nil {
			fmt.Println(err)
		}
		type PageData struct {
			Name      string
			Surname   string
			Email     string
			UserRole  string
			Dashboard domain.Dashboard
			IsExist   string
		}
		data := PageData{user.Name, user.Surname, email, userRole, DashBoard, isExistString}

		mainFile := "index.html"

		ts, err := template.ParseFiles(util.RolePageRender(app, userRole, mainFile)...)
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
