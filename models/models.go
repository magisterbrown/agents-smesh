package models

import (
    "database/sql"
)

var DB *sql.DB

type Player struct {
    Id int
    Name string
    Raiting float32
}

func GetLeaderboard() []Player {
    rows, _:= DB.Query("SELECT * from players")
    defer rows.Close()
    var res []Player
    for rows.Next() {
        var player Player
        rows.Scan(&player.Id, &player.Name, &player.Raiting);
        res = append(res, player);
    }
    return res
}
