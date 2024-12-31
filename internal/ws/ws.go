package ws

import (
	"log"
	"net/http"
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
	Receiver string `json:"receiver"`
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
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Println("ReadJSON error:", err)
			break
		}

		payload.Conn = *conn
		wsChan <- payload
	}

	mu.Lock()
	delete(clients, *conn)
	mu.Unlock()

	updateUserStatusInDB(payload.Username, false)
	users := getUserList()
	response := WsJonResponse{
		Action:         "list_users",
		ConnectedUsers: users,
	}
	BroadCastToAll(response)
}

func ListenToWsChannel() {
	for {
		e := <-wsChan
		var response WsJonResponse

		switch e.Action {
		case "username":
			mu.Lock()
			clients[e.Conn] = e.Username
			mu.Unlock()
			updateUserStatusInDB(e.Username, true)
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			BroadCastToAll(response)

		case "left":
			mu.Lock()
			delete(clients, e.Conn)
			mu.Unlock()
			updateUserStatusInDB(e.Username, false)
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			BroadCastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
			BroadCastToAll(response)

		case "private":
			response.Action = "private"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
			sendPrivateMessage(response, e.Receiver)
		}
	}
}

func getUserList() []string {
	var users []string
	rows, err := db.DB.Query("SELECT username FROM User WHERE is_online = 1")
	if err != nil {
		fmt.Println("Error fetching users from database:", err)
		return users
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			fmt.Println("Error scanning username:", err)
			continue
		}
		users = append(users, username)
	}
	return users
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

func sendPrivateMessage(response WsJonResponse, receiver string) {
	mu.Lock()
	defer mu.Unlock()
	for client, username := range clients {
		if username == receiver {
			err := client.WriteJSON(response)
			if err != nil {
				log.Printf("WebSocket error for client %v: %v", client, err)
				_ = client.Close()
				delete(clients, client)
			}
			break
		}
	}
}

func updateUserStatusInDB(username string, online bool) {
	_, err := db.DB.Exec("UPDATE User SET is_online = ? WHERE username = ?", online, username)
	if err != nil {
		fmt.Println("Error updating user status in database:", err)
	}
}
