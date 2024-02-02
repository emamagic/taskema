-- +migrate Up
CREATE TABLE columns (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL,
    creator_user_id INT NOT NULL,
    priority INT DEFAULT 0,
    workspace_id INT NOT NULL,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE columns;