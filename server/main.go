package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/auth"
	"real-time-forum/config"
	"real-time-forum/handlers"
)

var addr = flag.String("addr", ":8081", "http server address")

func main() {
	flag.Parse()
	/* ------------------------------ initialize db ----------------------------- */
	db := config.InitDB() // initializing connection
	defer db.Close()

	// Define the handlers here, to use it in both the wsServer & the API
	userHandler := &handlers.UserHandler{Db: db}
	messageHandler := &handlers.MessageHandler{Db: db}
	postHandler := &handlers.PostHandler{Db: db}
	commentHandler := &handlers.CommentHandler{Db: db}

	/* ---------- create websocket server and pass db connections to it --------- */
	wsServer := NewWebsocketServer(userHandler, messageHandler, postHandler, commentHandler)

	/* ------------------------ start chat server running ----------------------- */
	go wsServer.Run()

	api := &API{UserHandler: userHandler}
	middle := &auth.Middleware{UserHandler: userHandler}

	/* ------------- upgrade http conn and listen to client requests ------------ */
	http.HandleFunc("/ws", middle.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	}))
	/* ------------------------- api endpoints for auth ------------------------- */
	http.HandleFunc("/api/login", api.HandleLogin)
	http.HandleFunc("/api/signup", api.HandleSignup)

	fmt.Println("Hello world! The server is ready to serve!")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
