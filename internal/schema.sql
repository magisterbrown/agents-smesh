CREATE TABLE players (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(128) UNIQUE NOT NULL,
    --token VARCHAR(255) NOT NULL,
    raiting FLOAT NOT NULL
);

INSERT INTO players (user_name, raiting) VALUES ("brownie", 600),
    ("turbo", 1400),
    ("jack", 200);
