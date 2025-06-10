package api

import "net/http"

func WSDLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	http.ServeFile(w, r, "./wsdl/service.wsdl")
}
