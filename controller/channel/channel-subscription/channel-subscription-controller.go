package channel_subscription

import (
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"timtubeWebAdmin/config"
	"timtubeWebAdmin/controller/util"
	"timtubeWebAdmin/domain"
	user "timtubeWebAdmin/io/channel/channel"
	dcs "timtubeWebAdmin/io/channel/channel-subscription"
	channel_type "timtubeWebAdmin/io/channel/channel-type"
	user3 "timtubeWebAdmin/io/user/role"
	user2 "timtubeWebAdmin/io/user/user"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Get("/{id}", getHandler(app))
	r.Get("/get/{id}", getChannelHandler(app))
	r.Get("/user/{id}", getByUserHandler(app))
	r.Get("/channel/{id}", getByChannelHandler(app))
	r.Post("/create", createHandler(app))
	r.Get("/delete/{id}", deleteHandler(app))
	r.Post("/update", updateHandler(app))
	return r
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
			http.Redirect(w, r, "/channel/subscription", 301)
			return
		}
		channels, err := dcs.ReadChannelSubscription(id)
		if err != nil {
			fmt.Println(err, "error reading channelSubscription")
			http.Redirect(w, r, "/channel/subscription", 301)
			return
		}

		client, err := user2.ReadUser(channels.UserId)
		if err != nil {
			fmt.Println(err, "error reading client user")
		}
		channel, err := user.ReadChannel(channels.ChannelId)
		if err != nil {
			fmt.Println(err, "error reading channel")
		}
		owner, err := user2.ReadUser(channel.UserId)
		if err != nil {
			fmt.Println(err, "error reading owner user")
		}
		channelType, err := channel_type.ReadChannelType(channel.ChannelTypeId)
		if err != nil {
			fmt.Println(err, "error reading channeltype")
		}
		ownerRole, err := user3.ReadRole(owner.RoleId)
		if err != nil {
			fmt.Println(err, "error reading role")
		}
		clientRole, err := user3.ReadRole(client.RoleId)
		if err != nil {
			fmt.Println(err, "error reading role")
		}
		sEnc := base64.StdEncoding.EncodeToString(channel.Image)
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message     string
			Name        string
			Surname     string
			UserRole    string
			Channels    domain.ChannelSubscription
			Owner       domain.User
			OwnerRole   domain.Role
			Client      domain.User
			ClientRole  domain.Role
			Channel     domain.Channel
			ChannelType domain.ChannelType
			Image       string
			Email       string
		}
		data := PageData{message, name,
			surname, role,
			channels, owner, ownerRole, client,
			clientRole, channel, channelType, sEnc, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel-subscription/channel-subscription.html")...)
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

func getByUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", http.StatusTemporaryRedirect)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/", http.StatusTemporaryRedirect)
			return
		}
		channels, err := dcs.ReadChannelSubscriptionByUserId(id)
		if err != nil {
			fmt.Println(err, "error reading channels")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.ChannelSubscription
		}
		data := PageData{message, name, surname, role, channels}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-subscriptions.html")...)
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
func getByChannelHandler(app *config.Env) http.HandlerFunc {
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
		channels, err := dcs.ReadChannelSubscriptionByChannel(id)
		if err != nil {
			fmt.Println(err, "error reading channels")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.ChannelSubscription
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel/channel-subscriptions.html")...)
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
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}
		if id == "" {
			fmt.Println("wrong channel id")
			http.Redirect(w, r, "/channel/", 301)
			return
		}
		channels, err := GetRealChannelSubscription(id)
		if err != nil {
			fmt.Println(err, "error reading channels")
		}
		message := util.GetSession("message", app, r)
		type PageData struct {
			Message  string
			Name     string
			Surname  string
			UserRole string
			Channels []domain.ChannelSubscription
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel-subscription/channel-subscriptions.html")...)
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
			_, err := dcs.DeleteChannel(id)
			if err != nil {
				fmt.Println(err, "error creating channel")
			}
			app.Session.Put(r.Context(), "message", "You have Delete a new Channel")
			http.Redirect(w, r, "/channel/subscription", http.StatusTemporaryRedirect)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not Delete, data missing!")
		http.Redirect(w, r, "/channel/subscription", http.StatusTemporaryRedirect)
		return
	}
}
func createHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//Id        :=r.PostFormValue("id")
		ChannelId := r.PostFormValue("channelId")
		UserId := r.PostFormValue("user")
		Date := util.GetCurrentDate()

		if ChannelId != "" && UserId != "" {
			channelSubscriptionObject := domain.ChannelSubscription{"", ChannelId, UserId, Date}
			_, err := dcs.CreateChannelSubscription(channelSubscriptionObject)
			if err != nil {
				fmt.Println(err, "error creating channel Subscription")
			}
			app.Session.Put(r.Context(), "message", "You have created a new Subscription")
			http.Redirect(w, r, "/channel/subscription", http.StatusTemporaryRedirect)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/channel/subscription", http.StatusTemporaryRedirect)
		return
	}
}
func updateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		Id := r.PostFormValue("id")
		ChannelId := r.PostFormValue("channelId")
		UserId := r.PostFormValue("user")
		Date := util.GetCurrentDate()

		if Id != "" && ChannelId != "" && UserId != "" {
			channelSubscriptionObject := domain.ChannelSubscription{Id, ChannelId, UserId, Date}
			_, err := dcs.UpdateChannelSubscription(channelSubscriptionObject)
			if err != nil {
				fmt.Println(err, "error update channel Subscription")
			}
			app.Session.Put(r.Context(), "message", "You have update a new Subscription")
			http.Redirect(w, r, "/channel/subscription", http.StatusTemporaryRedirect)
			return
		}
		app.Session.Put(r.Context(), "message", "Could not create, data missing!")
		http.Redirect(w, r, "/channel/subscription", http.StatusTemporaryRedirect)
		return
	}
}
func GetChannelSubscription(channelSubscriptions []domain.ChannelSubscription) {

}
func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email, name, surname, role := util.GetPermenentSession(app, r)
		if email == "" {
			http.Redirect(w, r, "/user/account/login", 301)
			return
		}

		channels, err := user.ReadChannels()
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
			Email    string
		}
		data := PageData{message, name, surname, role, channels, email}

		ts, err := template.ParseFiles(util.RolePageRender(app, role, "channel-subscription/channels.html")...)
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
func getSubscriptions() ([]domain.ChannelSubscription, error) {
	var subscriptions []domain.ChannelSubscription
	channels, err := dcs.ReadChannelSubscriptions()
	if err != nil {
		fmt.Println(err, "error reading channels")
		return subscriptions, err
	}
	for _, subscription := range channels {
		channel, err := user.ReadChannel(subscription.ChannelId)
		if err != nil {
			fmt.Println(err, " error reading Channel.")
		} else {
			subscriptions = append(subscriptions, domain.ChannelSubscription{subscription.Id, channel.Name, subscription.UserId, subscription.Date})
			channel = domain.Channel{}
		}
	}
	return subscriptions, nil
}
