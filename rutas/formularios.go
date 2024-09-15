package rutas

import (
	"html/template"
	"io"
	"main/utils"
	"main/validaciones"
	"net/http"
	"os"
	"strings"
	"time"
)

func FormulariosGet(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/formulario.html", utils.Frontend))
	css_session, css_mensaje := utils.RetornarMensajesFlash(w, r)
	data := map[string]string{
		"css":     css_session,
		"mensaje": css_mensaje,
	}
	template.Execute(w, data)
}
func FormulariosPost(w http.ResponseWriter, r *http.Request) {
	mensaje := ""
	if len(r.FormValue("nombre")) == 0 {
		mensaje = mensaje + "el campo nombre está vacio\n"
	}
	if len(r.FormValue("correo")) == 0 {
		mensaje = mensaje + "el campo E-mail está vacio"
	}
	if validaciones.Regex_correo.FindAllStringSubmatch(r.FormValue("correo"), -1) == nil {
		mensaje = mensaje + "el E-mail ingresado no es valido"
	}
	if !validaciones.ValidarPassword(r.FormValue("password")) {
		mensaje = mensaje + "La contraseña debe tener almenos 1 numero, una mayuscula y una extension entre 6 y 20 caracteres"
	}
	if mensaje != "" {
		//fmt.Fprintln(w, mensaje)
		//return
		utils.CrearMensajesFlash(w, r, "danger", mensaje)
		http.Redirect(w, r, "/formularios", http.StatusSeeOther)
	}
}

func FormulariosUploadGet(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/upload.html", utils.Frontend))
	css_session, css_mensaje := utils.RetornarMensajesFlash(w, r)
	data := map[string]string{
		"css":     css_session,
		"mensaje": css_mensaje,
	}
	template.Execute(w, data)
}

func FormulariosUploadPost(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("foto")
	if err != nil {
		utils.CrearMensajesFlash(w, r, "danger", "ocurrió un error inesperado")
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "public/uploads/fotos/" + foto
	f, errCopy := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
	if errCopy != nil {
		utils.CrearMensajesFlash(w, r, "danger", "ocurrió un error inesperado")
	}
	_, errCopiar := io.Copy(f, file)
	if errCopiar != nil {
		utils.CrearMensajesFlash(w, r, "danger", "ocurrió un error inesperado")
	}
	utils.CrearMensajesFlash(w, r, "success", "Se subió el archivo "+foto+" exitosamente")
	http.Redirect(w, r, "/formularios/upload", http.StatusSeeOther)
}
