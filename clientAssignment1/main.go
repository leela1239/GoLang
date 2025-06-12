package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/greet/{user_name}", func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("user_name")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello " + username))

	})

	
	router.HandleFunc("/add-user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		userid := r.PathValue("user_id")
		w.Write([]byte("UserId : " + userid + "is added"))
	})

	router.HandleFunc("/get-user-details", func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query().Get("user_id")

		userDetails := map[string]string{
			"user_id":  userid,
			"username": "Leela",
			"email":    "leela@gmail.com",
		}

		w.Header().Set("Content-Type", "application/json")
		jsonResponse, err := json.Marshal(userDetails)
		if err != nil {
			http.Error(w, "Error generating JSON", http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)
	})

	http.ListenAndServe(":8080", router)

}
