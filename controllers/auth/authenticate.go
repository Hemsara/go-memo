package controller

import (
	"fmt"
	"net/http"
)

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Authentication successful!")
}
