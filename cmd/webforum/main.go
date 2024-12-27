package main

import (
	"log"

	"github.com/eediallo/real_time_forum/internal/db"
	"github.com/eediallo/real_time_forum/internal/handler"
	"github.com/eediallo/real_time_forum/internal/servers"
)

func main() {

	errInitDB := db.InitDB()
	if errInitDB != nil {
		log.Printf("error initiating db --main()--InitDB()--%s:", errInitDB.Error())
		return
	}

	defer db.CloseDB()

	log.Println("Initializing all templates")
	handler.InitTemplates()

	log.Println("Start running server")
	server, errRunServer := servers.RunServer()
	if errRunServer != nil {
		// log.Println(errRunServer)
		log.Fatal(errRunServer)
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
