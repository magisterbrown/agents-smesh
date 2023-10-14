package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3" 
)

var DB *sql.DB
func main() {
    DB, _ = sql.Open("sqlite3", "/tmp/rankings.db") 
    res, _ :=  DB.Exec("INSERT INTO submissions (user_id, container_id, raiting) VALUES (?, ?, ?)", 1, "aaa", 400.0);
    idx, _ := res.LastInsertId()
    fmt.Println(idx);
}
