package main

import (
	"fmt"
	"log"
	"net/http"

	// migration "github.com/awaisniaz/user-management/pkg/db/migration"
	"github.com/awaisniaz/user-management/internal/handlers"
	connection "github.com/awaisniaz/user-management/pkg/db"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Server is running on PORT: 3000")
	// migration.Up()
	connection.InitConnection()
	r := mux.NewRouter()

	r.HandleFunc("/registration", handlers.RegisterHandler)
	r.HandleFunc("/login",handlers.LoginHandler)
	log.Fatal(http.ListenAndServe(":8000", r))

}
