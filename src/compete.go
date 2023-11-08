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
    "encoding/json"
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

func startContainer(image string, command string) (types.HijackedResponse, string, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    errmsg := types.HijackedResponse{}
    if err != nil {
	    return errmsg, "", err
	}
    conf := container.Config{Image: image, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true, Cmd:  strslice.StrSlice{command}}
    playercont, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "")
    if err != nil {
	    return errmsg, "", err
	}
    hijack, err := cli.ContainerAttach(context.Background(), playercont.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
    if err != nil {
	    return errmsg, "", err
	}
    err = cli.ContainerStart(context.Background(), playercont.ID, types.ContainerStartOptions{})
    if err != nil {
	    return errmsg, "", err
	}
    if err != nil {
	    return errmsg, "", err
	}

    return hijack, playercont.ID, nil
}


func Match(player1 *models.Agent, player2 *models.Agent) error {
    
    cli, err := client.NewClientWithOpts(client.FromEnv)
    _ = cli
    if err != nil {
        fmt.Println(err)
    }

    agents := map[string]*models.Agent{
        "player_0": player1,
        "player_1": player2,
    }
    _ = agents


    //Start game container
    conf := container.Config{Image: config.GameTag, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true}
    gamecont, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "")
    if err != nil {
        panic(err)
    }
    hijack, err := cli.ContainerAttach(context.Background(), gamecont.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
    _ = hijack
    if err != nil {
        panic(err)
    }
    err = cli.ContainerStart(context.Background(), gamecont.ID, types.ContainerStartOptions{})
    if err != nil {
        panic(err)
    }

    //Play game
    fmt.Println(player1.Image)
    var message map[string]interface{}
    gameLoop:
    for {
        buf, _, err := hijack.Reader.ReadLine()
        err = json.Unmarshal(buf,&message)
        if err != nil{
            return err
        }
        intype, _ := message["type"].(string)
        switch intype {
            case "move":{
                agent_idx, _ := message["agent"].(string)
                output, containerID, err := startContainer(agents[agent_idx].Image, string(buf))
                if err != nil{
                    return err
                }
                line, _ , err := output.Reader.ReadLine()
                output.Close()
                err = cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: true})
                if err != nil{
                    return err
                }
                hijack.Conn.Write(append(line,[]byte("\n")...))
            }
            case "result":{
                break gameLoop
            }
            default:
                //Do nothing, but json is expected
        }
    }

    //Cleanup game
    err = cli.ContainerRemove(context.Background(), gamecont.ID, types.ContainerRemoveOptions{Force: true})
    if err != nil {
        panic(err)
    }
    fmt.Println("Played")
    return err


    }

