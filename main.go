package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "encoding/json"
    "ranking/models"
    _ "archive/zip"
    _ "github.com/mattn/go-sqlite3" 
    "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
    "context"


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
            file, _, err:= req.FormFile("submission")
            if err != nil {
                fmt.Printf(err.Error())
            	http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
            	return
            }
            defer file.Close()
            
            // Create docker image
	        cli, err := client.NewClientWithOpts(client.FromEnv)
            options := types.ImageBuildOptions{
                Tags: []string{"reformed"},
                SuppressOutput: true,                           
                Dockerfile: "submission/Dockerfile",           
            }                                                 
            var _ = options
            resp, err := cli.ImageBuild(context.Background(), file, options)
            if err != nil {
		        fmt.Println("Error:", err)
		        return
	        }
            defer resp.Body.Close()

            //for _, z := range zipr.File {

            //    fmt.Printf("type: %T\n", z.FileHeader)
            //    fmt.Printf("type: %s\n", z.FileHeader.Name)
            //}
            //          fmt.Printf("FIle :\n%s", header.Filename)

            //fmt.Fprintf(w, "Posted Solution")
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
