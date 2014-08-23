package sync

import (
	"net/http"

	"controllers"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	http.Handle("/", r)
}
