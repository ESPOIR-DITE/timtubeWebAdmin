package user

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	user3 "timtubeWebAdmin/io/user/role"
	user2 "timtubeWebAdmin/io/user/user"
	"timtubeWebAdmin/io/user/user-account"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", viewUsers(app))
	r.Get("/agent", viewAgents(app))
	r.Get("/admin", viewAdmins(app))
	r.Get("/All-users", viewAllUsers(app))
	r.Get("/users", viewUsers(app))
	r.Get("/super-admin", viewSuperAdmins(app))
	r.Get("/user/{email}", viewUser(app))
	r.Post("/user/update", updateUserDetails(app))

	return r
}

func viewAllUsers(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		user, err := user2.ReadUser(email)
		fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		users, err := user2.ReadUsers()
		if err != nil {
			fmt.Println(err, " error reading users")
		}
		type PageData struct {
			Users   []domain.User
			Name    string
			Surname string
			Message string
		}
		data := PageData{users, name, surname, message}
		var files = []string{}
		filesAgent := []string{
			app.Path + "admin/users.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "admin/users.html",
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

func viewAgents(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		user, err := user2.ReadUser(email)
		fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		users, err := user2.ReadAgentUsers()
		if err != nil {
			fmt.Println(err, " error reading users")
		}
		type PageData struct {
			Users   []domain.User
			Name    string
			Surname string
			Message string
		}
		data := PageData{users, name, surname, message}
		var files = []string{}
		filesAgent := []string{
			app.Path + "admin/users.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "admin/users.html",
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

func viewAdmins(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		user, err := user2.ReadUser(email)
		fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		users, err := user2.ReadAdminUsers()
		if err != nil {
			fmt.Println(err, " error reading users")
		}
		type PageData struct {
			Users   []domain.User
			Name    string
			Surname string
			Message string
		}
		data := PageData{users, name, surname, message}
		var files = []string{}
		filesAgent := []string{
			app.Path + "admin/users.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "admin/users.html",
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

func viewSuperAdmins(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		user, err := user2.ReadUser(email)
		fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		users, err := user2.ReadSuperAdminUsers()
		if err != nil {
			fmt.Println(err, " error reading users")
		}
		type PageData struct {
			Users   []domain.User
			Name    string
			Surname string
			Message string
		}
		data := PageData{users, name, surname, message}
		var files = []string{}
		filesAgent := []string{
			app.Path + "admin/users.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "admin/users.html",
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

func updateUserDetails(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		birthdate := r.PostFormValue("birthdate")
		email := r.PostFormValue("email")
		name := r.PostFormValue("name")
		surname := r.PostFormValue("surname")
		roleId := r.PostFormValue("roleId")
		if name != "" && email != "" && surname != "" && roleId != "" {
			userObject := domain.User{email, name, surname, birthdate, roleId}
			_, err := user2.UpdateUser(userObject)
			if err != nil {
				fmt.Println(err, " error update user")
				util.CreateSession("message", "couldn't update User details try again later", app, r)
				http.Redirect(w, r, "/user/user/user/"+email, 301)
				return
			}
		} else {
			util.CreateSession("message", "couldn't update User details try again later! required value missing.", app, r)
			http.Redirect(w, r, "/user/user/user/"+email, 301)
			return
		}
		util.CreateSession("message", "Successfully updated user Details.", app, r)
		http.Redirect(w, r, "/user/user/user/"+email, 301)
		return
	}
}

func viewUser(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		emailToview := chi.URLParam(r, "email")
		email, name, surname, _ := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		userToview, err := user2.ReadUser(emailToview)
		if err != nil {
			fmt.Println(err, " error reading user!")
		}
		userAccount, err := user.ReadUserAccountWithEmail(emailToview)
		if err != nil {
			fmt.Println(err, " error reading user Account!")
		}
		fmt.Println(userAccount)
		role, err := util.GetRoleName(userToview.RoleId)
		if err != nil {
			fmt.Println(err, " error reading user Role!")
		}
		roles, err := user3.ReadRoles()
		if err != nil {
			fmt.Println(err, " error reading roles")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Name        string
			Surname     string
			User        domain.User
			UserAccount domain.UserAccount
			Role        string
			Roles       []domain.Role
			Message     string
		}
		data := PageData{name, surname, userToview, userAccount, role, roles, message}
		files := []string{
			app.Path + "admin/user.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
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

func viewUsers(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, userRole := util.GetPermenentSession(app, r)
		message := util.GetSession("message", app, r)
		user, err := user2.ReadUser(email)
		fmt.Println(user)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		users, err := user2.ReadAllUsers()
		if err != nil {
			fmt.Println(err, " error reading users")
		}
		type PageData struct {
			Users   []domain.User
			Name    string
			Surname string
			Message string
		}
		data := PageData{users, name, surname, message}
		var files = []string{}
		filesAgent := []string{
			app.Path + "admin/users.html",
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		fileAdmins := []string{
			app.Path + "admin/users.html",
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
