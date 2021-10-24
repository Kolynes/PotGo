package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	connection *websocket.Conn
}

func (controller *WebSocketController) Upgrade(writer http.ResponseWriter, request *http.Request) (bool, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }

	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return false, err
	}
	controller.connection = connection

	controller.reader()
	connection.SetCloseHandler(controller.onClose)
	return true, nil
}

func (controller *WebSocketController) reader() {
	for {
		messageType, text, err := controller.connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		controller.onMessage(messageType, text)
	}
}

func (controller *WebSocketController) onClose(code int, text string) error {
	return nil
}

func (controller *WebSocketController) onMessage(messageType int, text []byte) {}
