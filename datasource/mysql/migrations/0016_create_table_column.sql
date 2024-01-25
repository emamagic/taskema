-- +migrate Up
CREATE TABLE columns (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL,
    priority INT DEFAULT 0,
    board_id INT NOT NULL,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (board_id) REFERENCES board(id) ON DELETE CASCADE,
);

-- +migrate Down
DROP TABLE columns;