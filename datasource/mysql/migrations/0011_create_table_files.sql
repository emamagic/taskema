-- +migrate Up
CREATE TABLE files (
    id INT PRIMARY KEY AUTO_INCREMENT,
    hash VARCHAR(191) NOT NULL UNIQUE,
    path VARCHAR(191) NOT NULL,
    user_creator_id INT NOT NULL,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_creator_id) REFERENCES users(id)
);

-- +migrate Down
DROP TABLE files;