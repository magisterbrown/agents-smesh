package main

import (
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
    "context"
    "fmt"
    "encoding/json"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
    file, err := archive.TarWithOptions("games/", &archive.TarOptions{IncludeFiles: []string{"Dockerfile", "test.py","play.py", "requirements.txt"}})
    var tag = "game:latest"
    options := types.ImageBuildOptions{
        Tags: []string{tag},
        SuppressOutput: true,                           
        Dockerfile: "Dockerfile",           
    }
    resp, err := cli.ImageBuild(context.Background(), file, options)
    if err != nil {
	    fmt.Println("Error:", err)
	    return
	}
    //Starter 
    conf := container.Config{Image: tag, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true}
    container, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "")
    hijack, err := cli.ContainerAttach(context.Background(), container.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
    cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
    if err != nil {
	    fmt.Println("Error:", err)
	    return
	}
    //Play
    //TODO: react to random commands with skip
    buf := make([]byte, 1024)
    var message map[string]interface{}
    var sz int
    movesLoop:
    for i:=0; i<6; i++ {
        sz, err = hijack.Conn.Read(buf)
        err = json.Unmarshal(buf[:sz], &message)
        if err != nil{
            fmt.Println("Game died") 
            return 
        }
        intype, _ := message["type"].(string)
        fmt.Println(intype)
        switch intype {
            case "move": {
                fmt.Printf(intype)
                //hijack.Conn.Write([]byte("1\n"))
                hijack.Conn.Write(append([]byte(`{"type":"decision", "choice":1}`), []byte("\n")...))

                //hijack.Conn.Read(buf)
                for key, value := range message{
                    fmt.Printf("%s: %v\n", key, value)
                }
            }
            case "result": {
                    fmt.Printf("WInner")
                    fmt.Printf(string(buf[:sz]))
                    break movesLoop
                }
            default: {
                    //fmt.Printf("Errror out")
                    //break movesLoop
            }
        }
    }

    defer hijack.Close()
    //cli.ContainerExecStart(context.Background(), execResp.ID, types.ExecStartCheck{})
    //resp, err := cli.ContainerExecCreate(context.Background(), file, options)
    //err2 := cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{}) 
    //if err2 != nil {
	//    fmt.Println("Error:", err2)
	//    return
	//}
    if err != nil {
	    fmt.Println("Error:", err)
	    return
	}
    defer resp.Body.Close()
}
