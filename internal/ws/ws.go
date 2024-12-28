package ws

import (
	"log"
	"net/http"
	"sort"
	"sync"

	"fmt"

	"github.com/eediallo/real_time_forum/internal/db"
	"github.com/gorilla/websocket"
)

var (
	wsChan  = make(chan WsPayload)
	clients = make(map[WebSocketConnection]string)
	mu      sync.Mutex
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WsJonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
	User     db.User
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	log.Println("Client connected to endpoint")

	conn := WebSocketConnection{Conn: ws}
	mu.Lock()
	clients[conn] = ""
	mu.Unlock()

	response := WsJonResponse{Message: `<em><small>Connected to server</small><em>`}
	err = ws.WriteJSON(response)
	if err != nil {
		log.Println("WriteJSON error:", err)
		return
	}

	go ListenForWs(&conn)
}

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", fmt.Sprintf("%v", r))
		}
		mu.Lock()
		delete(clients, *conn)
		mu.Unlock()
		conn.Close()
	}()

	for {
		var payload WsPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Printf("ReadJSON error for client %v: %v", conn, err)
			return
		}
		payload.Conn = *conn
		wsChan <- payload
	}
}

func ListenToWsChannel() {
	for {
		e := <-wsChan
		var response WsJonResponse

		switch e.Action {
		case "username":
			mu.Lock()
			clients[e.Conn] = e.User.Username
			users := getUserList()
			mu.Unlock()
			response.Action = "list_users"
			response.ConnectedUsers = users
			BroadCastToAll(response)

		case "left":
			mu.Lock()
			delete(clients, e.Conn)
			users := getUserList()
			mu.Unlock()
			response.Action = "list_users"
			response.ConnectedUsers = users
			BroadCastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.User.Username, e.Message)
			BroadCastToAll(response)
		}
	}
}

func getUserList() []string {
	mu.Lock()
	defer mu.Unlock()
	var userList []string
	for _, username := range clients {
		if username != "" {
			userList = append(userList, username)
		}
	}
	sort.Strings(userList)
	return userList
}

func BroadCastToAll(response WsJonResponse) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Printf("WebSocket error for client %v: %v", client, err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}
