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
