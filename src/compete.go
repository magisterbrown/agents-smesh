package compete

import (
    "ranking/models"
    "os/exec"
    "fmt"
    "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
    "context"
    "io"
)

func InitGame() error {
    file, err := archive.TarWithOptions("games/", &archive.TarOptions{IncludeFiles: []string{"Dockerfile", "test.py","play.py", "requirements.txt"}})
    if err != nil {
	    return err
	}
    var tag = "game:latest"
    

    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
	    return err
	}

    options := types.ImageBuildOptions{
        Tags: []string{tag},
        SuppressOutput: true,                           
        Dockerfile: "Dockerfile",           
    }
    _,_,_ = options, file,cli
    resp, err := cli.ImageBuild(context.Background(), file, options)
    if b, err := io.ReadAll(resp.Body); err == nil {
            fmt.Println(string(b))
        }
    _ = resp
    //defer resp.Body.Close()

    return err
}

func Match(player1 *models.Agent, player2 *models.Agent) (float32, error) {
    fmt.Println("GAAME")
    cmd := exec.Command("python3", "games/test.py")
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    return 1.0, nil
}

