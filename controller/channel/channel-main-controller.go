package user

import (
	"github.com/go-chi/chi"
	"net/http"
	"timtubeWebAdmin/config"
	cc "timtubeWebAdmin/controller/channel/channel"
	csc "timtubeWebAdmin/controller/channel/channel-subscription"
	ctc "timtubeWebAdmin/controller/channel/channel-type"
	cvc "timtubeWebAdmin/controller/channel/channel-video"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Mount("/", cc.Home(app))
	mux.Mount("/subscription", csc.Home(app))
	mux.Mount("/type", ctc.Home(app))
	mux.Mount("/video", cvc.Home(app))
	return mux
}
