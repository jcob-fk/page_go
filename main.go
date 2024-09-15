package main

import (
	"fmt"
	"log"
	"main/rutas"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros)
	mux.HandleFunc("/parametrosquery", rutas.ParametrosQueryString)
	mux.HandleFunc("/estructuras", rutas.Estructuras)

	mux.HandleFunc("/formularios", rutas.FormulariosGet)
	mux.HandleFunc("/formularios-post", rutas.FormulariosPost).Methods("POST")

	mux.HandleFunc("/formularios/upload", rutas.FormulariosUploadGet)
	mux.HandleFunc("/formularios/uploadPost", rutas.FormulariosUploadPost).Methods("POST")

	//Archivos estaticos hacia mux
	s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	mux.PathPrefix("/assets/").Handler(s)

	// Error 404
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(rutas.Pagina404).GetHandler()

	//Ejecucion de servidor
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	server := &http.Server{
		Addr:         "localhost:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("corriendo servidor desde http://localhost:" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}

/*
func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "hola mundo")
	})
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
*/
