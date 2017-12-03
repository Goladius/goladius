package Config

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, funcMap template.FuncMap, fromWhere string, file string, data interface{}, components []string) {
	t, err := template.New("").Funcs(funcMap).ParseFiles("resources/Views/"+fromWhere+"/"+file+".html", "resources/Views/"+fromWhere+"/layouts/base.html")

	ComponentsRender(t, fromWhere, components)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "main", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ComponentsRender(t *template.Template, fromWhere string, components []string) {
	for _, component := range components {
		t.ParseFiles("resources/Views/" + fromWhere + "/components/" + component + ".html")
	}

}

func RenderErrorTemplate(w http.ResponseWriter, fromWhere string, file string, data interface{}) {
	t, err := template.New("").ParseFiles("resources/Views/"+fromWhere+"/"+file+".html", "resources/Views/"+fromWhere+"/layouts/base.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "main", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
