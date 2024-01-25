-- +migrate Up
CREATE TABLE organizations(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL,
    avatar TEXT NULL,
    creator_user_id INT NOT NULL,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (creator_user_id) REFERENCES users(id)  ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE organizations;