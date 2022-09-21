package user_account

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	user_details "timtubeWebAdmin/domain/user/user-details"
	user2 "timtubeWebAdmin/io/user/user"
	"timtubeWebAdmin/io/user/user-account"
	user3 "timtubeWebAdmin/io/user/user-bank"
	user4 "timtubeWebAdmin/io/user/user-detail"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/login", loginHandler(app))
	r.Get("/forget", forgetHandler(app))
	r.Post("/login", loginPostHandler(app))
	r.Get("/register", registerHandler(app))
	r.Post("/register", registerPostHandler(app))
	r.Post("/update", updateUserAccount(app))
	r.Post("/user-details", createUserDetails(app))

	return r
}

func createUserDetails(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		crNumber := r.PostFormValue("crNumber")
		Tnumber := r.PostFormValue("Tnumber")
		bankType := r.PostFormValue("bankType")
		bankName := r.PostFormValue("bankName")
		email := r.PostFormValue("email")
		accountNumber := r.PostFormValue("accountNumber")
		branchCode := r.PostFormValue("branchCode")
		ccv := r.PostFormValue("ccv")

		if crNumber != "" && Tnumber != "" && bankType != "" && bankName != "" && accountNumber != "" && ccv != "" && branchCode != "" {
			userBankObject := user_details.UserBank{"", email, bankType, bankName, branchCode, accountNumber, ccv}
			userBank, err := user3.CreateUserBank(userBankObject)
			if err != nil {
				util.CreateSession("message", "Error occurred, try again.", app, r)
				fmt.Println(err, "error creating UserBank")
			} else if userBank.Id != "" {
				userDetails := user_details.UserDetails{"", email, userBank.Id, crNumber, Tnumber}
				_, err := user4.CreateUserDetails(userDetails)
				if err != nil {
					util.CreateSession("message", "Error occurred, try again.", app, r)
					_, err = user3.DeleteUserBankById(userBank.Id)
					if err != nil {
						fmt.Println(err, " error deleting User bank!")
					}
				}
			}
			util.CreateSession("message", "Successfully create your Bank details", app, r)

		} else {
			util.CreateSession("message", "Error occurred, try again.", app, r)
		}

		http.Redirect(w, r, "/", 301)
		return
	}
}

func updateUserAccount(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		customerId := r.PostFormValue("customerId")
		status, _ := strconv.ParseBool(r.PostFormValue("status"))
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		date := r.PostFormValue("birthdate")
		fmt.Println(date)
		if customerId != "" && email != "" && password != "" {
			//oldUserAccount, err := user.ReadUserAccount(customerId)
			//if err != nil {
			//	http.Redirect(w, r, "/user/user/", 301)
			//	return
			//}
			userAccount := domain.UserAccount{customerId, email, password, date, status, ""}
			fmt.Println("userAccount: ", userAccount)
			_, err := user.UpdateUserAccount(userAccount)
			if err != nil {
				fmt.Println(err, " error updating user")
				http.Redirect(w, r, "/user/user/user/"+email, 301)
				return
			}
		} else if email != "" {
			http.Redirect(w, r, "/user/user/user/"+email, 301)
			return
		}
		http.Redirect(w, r, "/user/user/", 301)
		return
	}

}

func forgetHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "user/forgot.html",
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

func registerPostHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		lastName := r.PostFormValue("lastName")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		if name != "" && lastName != "" && email != "" && password != "" {
			now := time.Now()
			fmt.Println(now.Format(util.YYYMMDD))
			userObject := domain.User{email, name, lastName, now.Format(util.YYYMMDD), ""}
			_, err := user2.CreateUser(userObject)
			if err != nil {
				fmt.Println(err, " error creating user")
				http.Redirect(w, r, "/user/account/register", 301)
				return
			}
			userAccount := domain.UserAccount{"", email, password, util.YYYMMDD, false, ""}
			_, err = user.CreateUserAccount(userAccount)
			if err != nil {
				user2.DeleteUser(email)
				fmt.Println(err, " error creating user")
				http.Redirect(w, r, "/user/account/register", 301)
				return
			}

		} else {
			http.Redirect(w, r, "/user/account/register", 301)
			return
		}
		http.Redirect(w, r, "/user/account/login", 301)
	}
}

func loginPostHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		//fmt.Println(email)
		//fmt.Println(password)
		if email != "" && password != "" {
			loginDetail := domain.UserAccount{"", email, password, "", false, ""}
			userAccount, err := user.LoginUserAccount(loginDetail)
			if err != nil {
				fmt.Println(err, " error login user")
				app.Session.Put(r.Context(), "message", "Check your login details and try again!")
				http.Redirect(w, r, "/user/account/login", 301)
				return
			}
			util.CreateSession("email", email, app, r)
			util.CreateSession("userId", userAccount.CustomerId, app, r)
			util.CreateSession("token", userAccount.Token, app, r)
			api.TOKEN = userAccount.Token
			http.Redirect(w, r, "/", 301)
			return
		} else {
			app.Session.Put(r.Context(), "message", "Check your login details there is missing!")
			http.Redirect(w, r, "/user/account/register", 301)
			return
		}
	}
}

func registerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message string
		}
		data := PageData{message}
		files := []string{
			app.Path + "user/register.html",
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

func loginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message string
		}
		data := PageData{message}
		files := []string{
			app.Path + "user/login.html",
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
