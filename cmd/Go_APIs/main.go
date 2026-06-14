package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/Signal"
	"syscall"

	"github.com/DrDead0/Go_APIs/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()


	// database setup

	//setup router
	router:= http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Welcome to test api"))
	})

	server:= http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	fmt.Println("Server Started...!")
	done := make(chan os.Signal,1)
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done


	//setup server
}