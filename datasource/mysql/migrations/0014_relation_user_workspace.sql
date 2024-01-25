-- +migrate Up
CREATE TABLE user_workspace (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    workspace_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE user_workspace;