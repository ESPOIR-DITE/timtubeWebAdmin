package util

import (
	"net/http"
	"time"
	"timtubeWebAdmin/config"
	user "timtubeWebAdmin/io/user/role"
)

const YYYMMDD = "2006-02-01"
const YYYYMMDDHHMMSS = "2006-01-02 15:04 05"

var BirthDate = time.Date(1990, time.November, 10, 23, 0, 0, 0, time.UTC)

func isEquals(int1, int2 int) bool {
	if int1 == int2 {
		return false
	}
	if int1 > int2 {
		return true
	}
	return false
}
func IsProcessing(fileCreationDate string) bool {
	now := time.Now()
	creationDate, _ := time.Parse(YYYYMMDDHHMMSS, fileCreationDate)
	myTime := creationDate.Add(time.Minute * 10)
	myTimeFormated, _ := time.Parse(YYYYMMDDHHMMSS, myTime.Format(YYYYMMDDHHMMSS))
	nowFormated, _ := time.Parse(YYYYMMDDHHMMSS, now.Format(YYYYMMDDHHMMSS))
	if myTimeFormated.Before(nowFormated) {
		return false
	}
	return true
}
func GetSession(key string, app *config.Env, r *http.Request) string {
	message := app.Session.GetString(r.Context(), key)
	if message != "" && key != "email" {
		app.Session.Remove(r.Context(), key)
	}
	return message
}
func GetRoleName(roleId string) (string, error) {
	role, err := user.ReadRole(roleId)
	return role.Name, err
}
func GetUserId(app *config.Env, r *http.Request) string {
	id := app.Session.GetString(r.Context(), "userId")
	return id
}

// email, name, surname, role
func GetPermenentSession(app *config.Env, r *http.Request) (string, string, string, string) {
	email := app.Session.GetString(r.Context(), "email")
	name := app.Session.GetString(r.Context(), "name")
	surname := app.Session.GetString(r.Context(), "surname")
	role := app.Session.GetString(r.Context(), "role")
	return email, name, surname, role
}
func CreateSession(key, value string, app *config.Env, r *http.Request) {
	app.Session.Remove(r.Context(), key)
	app.Session.Put(r.Context(), key, value)
}
func IsAdmin(email string) bool {
	return true
}

func GetDateTime(date string) time.Time {
	result, err := time.Parse(YYYMMDD, date)
	if err != nil {
		return time.Time{}
	}
	return result
}
