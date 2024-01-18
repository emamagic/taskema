-- +migrate Up
CREATE TABLE workspaces (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NULL,
    avatar TEXT NULL,
    creator_user_id INT NOT NULL,
    organization_id INT NOT NULL,
    priority INT DEFAULT 0,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE workspaces;