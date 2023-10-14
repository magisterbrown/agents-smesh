package models

import (
    "database/sql"
)

var DB *sql.DB

type Agent struct {
    Image string `json:"stream"`
    Raiting float32
}

func SaveAgent(data Agent, owner Player) (int64, error) {
    data.Raiting = 700.0
    res, err := DB.Exec("INSERT INTO submissions (user_id, container_id, raiting) VALUES (?, ?, ?)", owner.Id, data.Image, data.Raiting);
    id, _ := res.LastInsertId();
    return id, err
}

type Player struct {
    Id int
    Name string
    Password string
}

func GetPlayer(token string) (Player, error) {
    // TODO: defence from sql injections.
    var user Player
    err := DB.QueryRow("select * from players where id in (select user_id from sessions where token='"+token+"')").Scan(&user.Id, &user.Name, &user.Password);
    return user, err
}

func GetLeaderboard() []Player {
    rows, _:= DB.Query("SELECT * FROM players")
    defer rows.Close()
    var res []Player
    for rows.Next() {
        var player Player
        rows.Scan(&player.Id, &player.Name, &player.Password);
        res = append(res, player);
    }
    return res
}
