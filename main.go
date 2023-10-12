package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "encoding/json"
    "ranking/models"
    "archive/zip"
    _ "github.com/mattn/go-sqlite3" 

)

func getLeaderboard(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch req.Method {
        case "GET", "HEAD":
            res := models.GetLeaderboard()
            fmt.Fprintf(w, req.Method)
            //res := models.Player{Id: 0,
            //Name: "name",
            //Raiting: 12.3}
            json.NewEncoder(w).Encode(res)
        case "POST":
            req.ParseMultipartForm(11 << 20) 
            file, header, err:= req.FormFile("submission")
            if err != nil {
                fmt.Printf(err.Error())
            	http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
            	return
            }
            defer file.Close()


            fmt.Printf("type: %T\n", file)
            fmt.Printf("eee")
            zipr, _ := zip.NewReader(file, header.Size)
            for _, z := range zipr.File {

                fmt.Printf("type: %T\n", z.FileHeader)
                fmt.Printf("type: %s\n", z.FileHeader.Name)
            }
                      fmt.Printf("FIle :\n%s", header.Filename)

            fmt.Fprintf(w, "Posted Solution")
            json.NewEncoder(w).Encode(map[string]string{"status":"ok", "raiting": "600"})
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
