package compete

import (
    "ranking/models"
    "os/exec"
    "fmt"
)

func Match(player1 *models.Agent, player2 *models.Agent) (float32, error) {
    fmt.Println("GAAME")
    cmd := exec.Command("python3", "games/test.py")
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    return 1.0, nil
}
