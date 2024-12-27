package handler

import "net/http"

// NotFoundHandler handles non-existing routes
func NotFoundHandler(w http.ResponseWriter, r *http.Request) bool {
	validPaths := map[string]bool{
		"/":              true,
		"/users/sign_up": true,
		"/users/login":   true,
		"/dashboard":     true,
		"/users/logout":  true,
		"/profile":       true,
		"/post":          true,
		"/add_comment":   true,
	}

	if _, valid := validPaths[r.URL.Path]; !valid {
		w.WriteHeader(http.StatusNotFound)
		ErrorPageHandler(w, "404 - Page Not Found", nil)
		return false
	}

	return true
}
