package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3" 
)


func getLeaderboard(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "GET", "HEAD":
            fmt.Fprintf(w, req.Method)
        case "POST":
            fmt.Fprintf(w, "Posted Solution")
        default:
            w.WriteHeader(404)
    }
}

func postSolution(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        default:
            w.WriteHeader(404)
    }
}

var db *sql.DB

type Roww struct {
        ID     int64
}

func main() {
    db, err := sql.Open("sqlite3", "/tmp/infos.sql") 
    if err != nil {
        fmt.Print("Hello, ")
    }
    defer db.Close()

    // Query for all table names in the database
    rows, err := db.Query("SELECT * FROM basic;")

    // Iterate through the results and print table names
    for rows.Next() {
        var rrw Roww
        rows.Scan(&rrw.ID);
        fmt.Println("Table Name:", rrw)
    }
    defer rows.Close()
    fmt.Println("DONE:")


//	_, err = db.Exec(sqlStmt)
    //http.HandleFunc("/solution", postSolution)
    http.HandleFunc("/leaderboard", getLeaderboard)
    http.ListenAndServe(":8090", nil)
}
