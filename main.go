package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// create feedback in database
func createFeedback(w http.ResponseWriter, r *http.Request, db *sql.DB) {}

func main() {
	// figure out what port to listen on
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	// connect to database
	dbUrl := "postgresql://postgres@localhost:5432/feedback?sslmode=disable"
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ping database
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to database ❌")
	} else {
		log.Println("Connected to database ✅")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/api/feedback", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			createFeedback(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// start server
	log.Fatal(http.ListenAndServe(port, nil))
}
