package compete

import (
    "ranking/config"
    "ranking/models"
    "fmt"
    "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/api/types/strslice"
    "context"
    _ "encoding/json"
    "io/ioutil"
)

func InitGame() error {
    file, err := archive.TarWithOptions(config.GameFolder, &archive.TarOptions{IncludeFiles: []string{"Dockerfile", "test.py","play.py", "requirements.txt"}})
    if err != nil {
	    return err
	}
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
	    return err
	}

    options := types.ImageBuildOptions{
        Tags: []string{config.GameTag},
        SuppressOutput: true,                           
        Dockerfile: "Dockerfile",           
    }
    resp, err := cli.ImageBuild(context.Background(), file, options)
    defer resp.Body.Close()

    return err
}

func makeStep(player *models.Agent, world string) (string, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    errmsg := ""
    if err != nil {
	    return errmsg, err
	}
    conf := container.Config{Image: player.Image, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true, Cmd:  strslice.StrSlice{world}}
    playercont, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "")
    if err != nil {
	    return errmsg, err
	}
    hijack, err := cli.ContainerAttach(context.Background(), playercont.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
    if err != nil {
	    return errmsg, err
	}
    err = cli.ContainerStart(context.Background(), playercont.ID, types.ContainerStartOptions{})
    if err != nil {
	    return errmsg, err
	}
    data, err := ioutil.ReadAll(hijack.Reader)
    if err != nil {
	    return errmsg, err
	}

    return string(data), nil
}


func Match(player1 *models.Agent, player2 *models.Agent) error {
    
    cli, err := client.NewClientWithOpts(client.FromEnv)
    _ = cli
    output, err := makeStep(player1, `{"asus": 3}`)
    fmt.Println(output) 
    if err != nil {
        fmt.Println(err)
    }


    ////Start game container
    //conf := container.Config{Image: config.GameTag, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true}
    //gamecont, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "")
    //if err != nil {
    //    panic(err)
    //}
    //hijack, err := cli.ContainerAttach(context.Background(), gamecont.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
    //_ = hijack
    //if err != nil {
    //    panic(err)
    //}
    //err = cli.ContainerStart(context.Background(), gamecont.ID, types.ContainerStartOptions{})
    //if err != nil {
    //    panic(err)
    //}

    ////Play game
    //fmt.Println(player1.Image)
    //buf := make([]byte, 1024)
    //var message map[string]interface{}
    //var sz int
    //gameLoop:
    //for {
    //    sz, err = hijack.Conn.Read(buf)
    //    err = json.Unmarshal(buf[:sz], &message)
    //    if err != nil{
    //        return err
    //    }
    //    intype, _ := message["type"].(string)
    //    switch intype {
    //        case "move":{
    //            hijack.Conn.Write(append([]byte(`{"type":"decision", "choice":1}`), []byte("\n")...))
    //        }
    //        case "result":{
    //            break gameLoop
    //        }
    //        default:
    //            //Do nothing, but json is expected
    //    }
    //}

    ////Cleanup game
    //err = cli.ContainerRemove(context.Background(), gamecont.ID, types.ContainerRemoveOptions{Force: true})
    //if err != nil {
    //    panic(err)
    //}
    //fmt.Println("Played")
    return err


    }

