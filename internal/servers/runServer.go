package servers

import (
	"net/http"

	"github.com/eediallo/real_time_forum/internal/handler"
	"github.com/eediallo/real_time_forum/internal/middleware"
)

//
func RunServer() (*http.Server, error) {
	// available routes for all individuals on the internet
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomePage)
	mux.HandleFunc("/users/sign_up", handler.RegisterUser)
	mux.HandleFunc("/users/login", handler.LoginUser)
	mux.HandleFunc("/dashboard/", handler.DashboardPage)

	// Protected routes
	mux.Handle("/users/logout", middleware.AuthMiddleware(http.HandlerFunc(handler.LogoutUser)))
	mux.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(handler.ProfilePage)))
	mux.Handle("/post", middleware.AuthMiddleware(http.HandlerFunc(handler.PostHandler)))
	mux.Handle("/add_comment", middleware.AuthMiddleware(http.HandlerFunc(handler.AddComment)))
	mux.Handle("/like", middleware.AuthMiddleware(http.HandlerFunc(handler.LikePostHandler)))
	mux.Handle("/dislike", middleware.AuthMiddleware(http.HandlerFunc(handler.DislikePostHandler)))
	mux.Handle("/like_dislike_comment", middleware.AuthMiddleware(http.HandlerFunc(handler.LikeDislikeCommentHandler)))

	// Serve static files from the "static" directory
	// ????<<<<<<add custom function to handle error in cas dir not exist
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server, nil
}
