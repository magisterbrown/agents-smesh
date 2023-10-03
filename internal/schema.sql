
CREATE TABLE players (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_name VARCHAR(128) UNIQUE NOT NULL,
    raiting FLOAT NOT NULL
);
INSERT INTO players (id, user_name, raiting) VALUES 
    (3, "aa", 140.0),
    (4, "bb", 130.0),
    (5, "bc", 130.0);

