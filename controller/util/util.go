package util

import (
	"net/http"
	"time"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/domain"
	user "timtubeWebAdmin/io/user/role"
)

const YYYMMDD = "2006-01-02"
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
func GetCurrentDate() string {
	now := time.Now()
	return now.Format(YYYMMDD)
}
func GetDateTime(date string) time.Time {
	result, err := time.Parse(YYYMMDD, date)
	if err != nil {
		return time.Time{}
	}
	return result
}
func GetStringDateTime(date string) string {
	result, err := time.Parse(YYYMMDD, date)
	if err != nil {
		return time.Now().Format(YYYMMDD)
	}
	return result.Format(YYYMMDD)
}
func RolePageRender(app *config.Env, userRole string, pages string) []string {
	var filesAgent []string
	switch userRole {
	case "superAdmin":
		filesAgent = []string{
			app.Path + pages,
			app.Path + "template/super-admin-navigator.html",
			app.Path + "template/topbar.html",
		}
		break
	case "admin":
		filesAgent = []string{
			app.Path + pages,
			app.Path + "template/super-admin-navigator.html",
			app.Path + "template/topbar.html",
		}
		break
	case "user":
		filesAgent = []string{
			app.Path + pages,
			app.Path + "template/navigator.html",
			app.Path + "template/topbar.html",
		}
		break
	case "agent":
		filesAgent = []string{
			app.Path + pages,
			app.Path + "template/admin-navigator.html",
			app.Path + "template/topbar.html",
		}
		break
	}
	return filesAgent
}

type ChannelsVideosOwners struct {
	Channel     domain.Channel
	ChannelType domain.ChannelType
	User        domain.User
	Videos      domain.Video
}
