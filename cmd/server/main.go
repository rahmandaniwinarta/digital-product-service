package main

import (
	"fmt"
	"net/http"

	"digital-product-service/internal/database"
)

func main() {
	db := database.ConnectDatabase()
	database.DBMigrate(db, "up") //

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
