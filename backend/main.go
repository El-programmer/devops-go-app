package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
    var message string
    err := db.QueryRow("SELECT 'Connected to MySQL!'").Scan(&message)
    if err != nil {
        http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Hello from Go backend!\nDatabase says: %s", message)
}

func main() {
    var err error

    dsn := "user:password@tcp(mysql:3306)/mydb"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %s", err)
    }

    // Ping to ensure connection is alive
    if err := db.Ping(); err != nil {
        log.Fatalf("Database not reachable: %s", err)
    }

    http.HandleFunc("/", handler)
    fmt.Println("Server is running on port 8000")
    http.ListenAndServe(":8000", nil)
}
