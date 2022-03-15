package routes

import (
	"log"
	"net/http"

	"github.com/treinalinux/api-go-rest/controllers"
)

// HandleRequest on route root
func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
