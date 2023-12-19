package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,DELETE,PATCH,POST,PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Content-Type,X-CSRF-Token")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
