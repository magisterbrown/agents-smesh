package main

import (
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
    "context"
    "fmt"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
    file, err := archive.TarWithOptions("games/", &archive.TarOptions{IncludeFiles: []string{"Dockerfile", "test.py", "requirements.txt"}})
    var tag = "game:latest"
    options := types.ImageBuildOptions{
        Tags: []string{tag},
        SuppressOutput: true,                           
        Dockerfile: "Dockerfile",           
    }
    resp, err := cli.ImageBuild(context.Background(), file, options)
    //container, err := 
    conf := container.Config{Image: tag, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true}
    container, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "coolgame")
    fmt.Println(container.ID)
    hijack, err := cli.ContainerAttach(context.Background(), container.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
    cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
    //execResp, err := 
    if err != nil {
	    fmt.Println("Error:", err)
	    return
	}
    //io.Copy(os.Stdout, hijack.Reader)
    buf := make([]byte, 1024)
    sz, err := hijack.Conn.Read(buf)
    fmt.Printf(string(buf[:sz]))
    hijack.Conn.Write([]byte("Message received.\n"))
    sz, err = hijack.Conn.Read(buf)
    fmt.Printf(string(buf[:sz]))
    defer hijack.Close()
    //cli.ContainerExecStart(context.Background(), execResp.ID, types.ExecStartCheck{})
    //resp, err := cli.ContainerExecCreate(context.Background(), file, options)
    //err2 := cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{}) 
    //if err2 != nil {
	//    fmt.Println("Error:", err2)
	//    return
	//}
    defer resp.Body.Close()
}
