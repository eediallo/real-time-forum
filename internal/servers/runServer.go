package servers

import (
	"net/http"

	"github.com/eediallo/real_time_forum/internal/handler"
	"github.com/eediallo/real_time_forum/internal/middleware"
	"github.com/eediallo/real_time_forum/internal/ws"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// available routes for all individuals on the internet
	r.HandleFunc("/", handler.HomePage)
	r.HandleFunc("/users/sign_up", handler.RegisterUser)
	r.HandleFunc("/users/login", handler.LoginUser)

	// WebSocket route
	r.HandleFunc("/ws", ws.WsEndpoint)

	// Protected routes
	r.Handle("/users/logout", middleware.AuthMiddleware(http.HandlerFunc(handler.LogoutUser)))
	r.Handle("/users/{username}", middleware.AuthMiddleware(http.HandlerFunc(handler.ProfilePage))).Methods("GET")
	r.Handle("/post", middleware.AuthMiddleware(http.HandlerFunc(handler.PostHandler)))
	r.Handle("/add_comment", middleware.AuthMiddleware(http.HandlerFunc(handler.AddComment)))
	r.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(handler.DashboardPage))).Methods("GET")
	r.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(handler.PrivateChat))).Methods("POST")
	r.Handle("/like", middleware.AuthMiddleware(http.HandlerFunc(handler.LikePostHandler)))
	r.Handle("/dislike", middleware.AuthMiddleware(http.HandlerFunc(handler.DislikePostHandler)))
	r.Handle("/like_dislike_comment", middleware.AuthMiddleware(http.HandlerFunc(handler.LikeDislikeCommentHandler)))

	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	return r
}

func RunServer() (*http.Server, error) {
	r := SetupRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server, nil
}
