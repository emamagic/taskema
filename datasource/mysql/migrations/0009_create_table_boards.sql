-- +migrate Up
CREATE TABLE boards (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL,
    avatar TEXT NULL,
    creator_user_id INT NOT NULL,
    workspace_id INT NOT NULL,
    priority INT DEFAULT 0,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DELETE TABLE boards;