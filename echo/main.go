package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.StaticFile("/", "./index.html")
	r.GET("/ws", echoHandler)

	r.Use(cors.Default())

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

func echoHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for {
		t, data, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, data)
	}
}
