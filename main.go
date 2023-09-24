package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "ranking/models"

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

func main() {
    var err error
    models.DB, err = sql.Open("sqlite3", "/tmp/rankings.db") 
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/leaderboard", getLeaderboard)
    http.ListenAndServe(":8090", nil)
}
