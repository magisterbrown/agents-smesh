package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "encoding/json"
    "ranking/models"
    "ranking/src"
//    _ "archive/zip"
    _ "github.com/mattn/go-sqlite3" 
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "github.com/dustinkirkland/golang-petname"
    "context"
    "strings"
)

func getLeaderboard(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch req.Method {
        case "GET", "HEAD":
            res := models.GetLeaderboard()
            fmt.Fprintf(w, req.Method)
            json.NewEncoder(w).Encode(res)
        case "POST":
            req.ParseMultipartForm(11 << 20) 

            //Authentication
            player, err := models.GetPlayer(req.Header.Get("Authorization"));
            if err != nil {
            	http.Error(w, "Authorization failed", http.StatusUnauthorized)
                return 
            }

            //Processing form
            file, _, err:= req.FormFile("submission")
            if err != nil {
            	http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
            	return
            }
            defer file.Close()
            
            // Create docker image
	        cli, err := client.NewClientWithOpts(client.FromEnv)
            options := types.ImageBuildOptions{
                Tags: []string{petname.Generate(2, "-")+":player"},
                SuppressOutput: true,                           
                Dockerfile: "submission/Dockerfile",           
            }                                                 
            resp, err := cli.ImageBuild(context.Background(), file, options)


            
            if err != nil {
		        fmt.Println("Error:", err)
		        return
	        }
            defer resp.Body.Close()

            var submission models.Agent
            json.NewDecoder(resp.Body).Decode(&submission)
            submission.Image = strings.TrimSuffix(strings.TrimPrefix(submission.Image, "sha256:"), "\n")
            _, err = compete.Match(&submission, &submission);
            if(err != nil){
            	http.Error(w, "Agent does not play by the rules", http.StatusBadRequest)
            	return
            }
            subid, _ := models.SaveAgent(submission, player)
            fmt.Println(subid)
            // TODO: Schedule games

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

    err = compete.InitGame()

    if err != nil {
	    fmt.Println("Error:", err)
	    return
	}


    http.HandleFunc("/leaderboard", getLeaderboard)
    http.ListenAndServe(":8090", nil)
}
