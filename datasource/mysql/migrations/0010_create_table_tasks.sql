-- +migrate Up
CREATE TABLE tasks(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL,
    avatar TEXT NULL,
    creator_user_id INT NOT NULL,
    description TEXT,
    due_date BIGINT DEFAULT NULL,
    board_id INT NOT NULL,
    assigned_user_id INT DEFAULT 0,
    priority INT DEFAULT 0,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (creator_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (board_id) REFERENCES boards(id) ON DELETE CASCADE,
    FOREIGN KEY (assigned_user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- FOREIGN KEY (assigned_user_id) REFERENCES users(id)
-- +migrate Down
DELETE TABLE tasks;