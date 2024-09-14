package rutas

import (
	"html/template"
	"main/utils"
	"main/validaciones"
	"net/http"
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
