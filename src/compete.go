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
    if err != nil {
        panic(err)
    }
    conf := container.Config{Image: image, AttachStdin: true, AttachStdout: true, AttachStderr: true, Tty: true, OpenStdin: true, StdinOnce: true, Cmd:  strslice.StrSlice{command}}
    hijack := types.HijackedResponse{}
    playercontID := ""

    for range []int{1} {
        playercont, err := cli.ContainerCreate(context.Background(), &conf, nil, nil, nil, "")
        if err != nil {
            break
        }
        playercontID = playercont.ID
        hijack, err = cli.ContainerAttach(context.Background(), playercont.ID, types.ContainerAttachOptions{Stream:true, Stdout:true, Stdin:true, Stderr:true})
        if err != nil {
            break
        }
        if err = cli.ContainerStart(context.Background(), playercont.ID, types.ContainerStartOptions{}); err != nil {
            break
        }
    }

    return hijack, playercontID, nil
}


func Match(player1 *models.Agent, player2 *models.Agent) error {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
    
    //Start game container
    hijack, gamecontID, err := startContainer(config.GameTag, "")
    if err != nil {
        panic(err)
    }

    agents := map[string]*models.Agent{
            "player_0": player1,
            "player_1": player2,
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
    err = cli.ContainerRemove(context.Background(), gamecontID, types.ContainerRemoveOptions{Force: true})
    if err != nil {
        panic(err)
    }
    return err
    }
