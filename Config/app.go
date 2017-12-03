package Config

import (
	"net/http"

	"../resources/Langs"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
)

type app struct {
	AppName  string
	Language string
	Locale   string
}

func App() app {
	returnApp := app{
		AppName:  "Goladius",
		Language: "English",
		Locale:   "en",
	}
	return returnApp
}
// GetRoute
func GetRoute(router *mux.Router, routeName string, vals ...string) string {
	routeUrl, _ := router.Get(routeName).URL(vals...)
	return routeUrl.String()
}

// Lang Change Helper Func.
func ChangeLanguage(req *http.Request, res http.ResponseWriter, store sessions.Store) {
	lang := req.FormValue("lang")
	session, _ := store.Get(req, "goladius")
	// Set some session values.
	if lang == "en" {
		session.Values["language"] = "English"
	} else if lang == "tr" {
		session.Values["language"] = "Turkish"
	}
	session.Save(req, res)
}

var language string

// Translation func
func Translate(langSession string, data string) string {

	if langSession == "" {
		language = App().Language
	} else {
		language = langSession
	}

	Languages := map[string]func(data string) string{
		"English": Langs.English,
		"Turkish": Langs.Turkish,
	}

	return Languages[language](data)
}