package main

import (
	"SimpleServer/server"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	conn, err := pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s/%s", os.Args[1], os.Args[2], os.Args[3], os.Args[4]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	psqlServer := server.NewServer(conn)

	http.HandleFunc("/PSQL/JSON", psqlServer.PsqlHandler)

	fmt.Println("Starting....")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
