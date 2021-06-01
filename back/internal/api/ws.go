package api

import (
	"back/internal/models"
	"back/internal/services/jwt"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)


type message struct {
	RoomId uint64 `json:"room_id"`
	Body string `json:"body"`
}
type hub struct {
	connections map[*connection]bool
}

type connection struct {
	ws *websocket.Conn
	user *models.User
}


var h = hub{
	connections:      make(map[*connection]bool),
}



var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
}

func ServeWs(w http.ResponseWriter, r *http.Request)  {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return  true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	conn := &connection{
		ws:   ws,
	}

	user , err := jwt.GetUserFromToken(r.URL.Query().Get("token"))

	conn.user = user
	if err != nil {
		CloseAndRemoveConnection(conn)
		return
	}

	conn.ws.WriteJSON(gin.H{"message" :"Connected Successfully"})

	h.connections[conn] = true
	go func() {
		defer CloseAndRemoveConnection(conn)
		for {
			_, msg, err := conn.ws.ReadMessage()
			if err != nil {
				break
			}
			var messageFromSocket message
			err = json.Unmarshal(msg, &messageFromSocket)
			if err != nil {
				fmt.Println(err)
				break
			}

			insertMessage := models.Message{
				Body:       messageFromSocket.Body,
				FromUserId:  conn.user.Id,
				RoomId:     messageFromSocket.RoomId,
			}
			models.DB.Create(&insertMessage)
			insertMessage.FromUser = *conn.user
			for cons := range h.connections{
				err := cons.ws.WriteJSON(gin.H{
					"message" :insertMessage,
					"for_message": true,
				})
				if err != nil {
					fmt.Println(err)
					CloseAndRemoveConnection(conn)
				}
			}
		}
	}()
}

func CloseAndRemoveConnection(conn *connection)  {
	delete(h.connections , conn)
	conn.ws.Close()
}




