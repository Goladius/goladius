package Front

import (
	"net/http"

	"../../../Config"
	"github.com/jinzhu/gorm"
)

type HomePageData struct {
	Title  string
	Locale string
}

func HomeController(res http.ResponseWriter, req *http.Request, Db *gorm.DB) {

	data := HomePageData{
		Locale: Config.App().Locale,
		Title:  Config.App().AppName,
	}

	Config.RenderTemplate(res, nil, "", "index", data, nil)
}
