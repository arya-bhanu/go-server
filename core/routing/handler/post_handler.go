package handler

import (
	"fmt"
	"net/http"
)

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	jwtData := r.Context().Value("data").(map[string]any)
	fmt.Printf("%+v", jwtData["email"].(string))
	w.Write([]byte("Hello World"))
}
