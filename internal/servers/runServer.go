package servers

import (
	"net/http"

	"github.com/eediallo/real_time_forum/internal/handler"
	"github.com/eediallo/real_time_forum/internal/middleware"
	"github.com/gorilla/mux"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		// Read message from browser
// 		msgType, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			return
// 		}

// 		// Write message back to browser
// 		if err = conn.WriteMessage(msgType, msg); err != nil {
// 			return
// 		}
// 	}
// }

func RunServer() (*http.Server, error) {
	r := mux.NewRouter()

	// available routes for all individuals on the internet
	r.HandleFunc("/", handler.HomePage)
	r.HandleFunc("/users/sign_up", handler.RegisterUser)
	r.HandleFunc("/users/login", handler.LoginUser)
	r.HandleFunc("/dashboard", handler.DashboardPage)

	// WebSocket route
	// r.HandleFunc("/ws", wsEndpoint)

	// Protected routes
	r.Handle("/users/logout", middleware.AuthMiddleware(http.HandlerFunc(handler.LogoutUser)))
	r.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(handler.ProfilePage)))
	r.Handle("/post", middleware.AuthMiddleware(http.HandlerFunc(handler.PostHandler)))
	r.Handle("/add_comment", middleware.AuthMiddleware(http.HandlerFunc(handler.AddComment)))
	r.Handle("/like", middleware.AuthMiddleware(http.HandlerFunc(handler.LikePostHandler)))
	r.Handle("/dislike", middleware.AuthMiddleware(http.HandlerFunc(handler.DislikePostHandler)))
	r.Handle("/like_dislike_comment", middleware.AuthMiddleware(http.HandlerFunc(handler.LikeDislikeCommentHandler)))

	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server, nil
}
