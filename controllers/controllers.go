package controllers

import (
	"fmt"
	"net/http"
)

// Home this is homepage
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
