package main

import
(
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/FantasyApi/routers/handlers"
	"github.com/FantasyApi/conn"
)

func main()  {
	DB := conn.Conn()
	h := handlers.New(DB)
	router  := mux.NewRouter()


	router.HandleFunc("/createuser", h.CreateUser).Methods("POST")
}