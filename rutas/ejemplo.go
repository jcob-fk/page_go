package rutas

import (
	"html/template"
	"net/http"

	"main/utils"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("templates/home.html", utils.Frontend))
	template.Execute(w, nil)
	/*
		template, err := template.ParseFiles("templates/home.html", utils.Frontend)
		if err != nil {
			panic(err)
		} else {
			template.Execute(w, nil)
		}
	*/
}
func Pagina404(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("templates/404.html", utils.Frontend))
	template.Execute(w, nil)
	/*
		template, err := template.ParseFiles("templates/home.html", utils.Frontend)
		if err != nil {
			panic(err)
		} else {
			template.Execute(w, nil)
		}
	*/
}
func Nosotros(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/nosotros.html", utils.Frontend)
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}
func Parametros(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := map[string]string{
		"id":   vars["id"],
		"slug": vars["slug"],
	}
	template, err := template.ParseFiles("templates/parametros.html", utils.Frontend)
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)
	}
}
func ParametrosQueryString(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"id":   r.URL.Query().Get("id"),
		"slug": r.URL.Query().Get("slug"),
	}
	template, err := template.ParseFiles("templates/parametrosquery.html", utils.Frontend)
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)
	}
}

type Habilidad struct {
	Nombre string
}

type Dato struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/estructuras.html", utils.Frontend)
	habilidad1 := Habilidad{"Inteligencia"}
	habilidad2 := Habilidad{"Programacion"}
	habilidad3 := Habilidad{"videojuegos"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, Dato{"Camilo", 18, 1, habilidades})
	}
}

/*
func Nosotros(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola para todos nosotros el pepe")
}
func Parametros(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "ID = "+vars["id"]+" | SLUG = "+vars["slug"])
}
func ParametrosQueryString(w http.ResponseWriter, r *http.Request) {
	// Ruta completa: localhost:8081/parametrosquery?id=111&slug=mislug
	fmt.Fprintln(w, r.URL)                     // Imprime /parametrosquery?id=111&slug=mislug
	fmt.Fprintln(w, r.URL.RawQuery)            // Imprime id=111&slug=mislug
	fmt.Fprintln(w, r.URL.Query())             // Imprime map[id:[111] slug:[mislug]]
	fmt.Fprintln(w, r.URL.Query().Get("id"))   // Imprime 111
	fmt.Fprintln(w, r.URL.Query().Get("slug")) // Imprime mislug
	id := r.URL.Query().Get("id")
	slug := r.URL.Query().Get("slug")
	fmt.Fprintf(w, "id= %s | slug= %s", id, slug)
}
*/
