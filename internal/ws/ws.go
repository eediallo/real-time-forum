package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgradeConnection = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}


// WsJonResponse define the response sent back from the websocket
type WsJonResponse struct {
	Action string `json:"action"`
	Message string `json:"message"`
	MessageType string `json:"message_type"`
}


//WsEndpoint upgrade connection to websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request){
	 ws, err := upgradeConnection.Upgrade(w, r, nil)
	 if err != nil {
		log.Println(err)
	 }

	 log.Println("Client connected to endpoint")

	 var response WsJonResponse
	 response.Message = `<em><small>Connected to server</small><em>`

	 err = ws.WriteJSON(response)
	 if err != nil {
		log.Println(err)
	 }
}