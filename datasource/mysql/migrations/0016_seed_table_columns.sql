-- +migrate Up
INSERT INTO columns (title, workspace_id, creator_user_id) VALUES ("to-do", 1, 1);
INSERT INTO columns (title, workspace_id, creator_user_id) VALUES ("in-progress", 1, 1);
INSERT INTO columns (title, workspace_id, creator_user_id) VALUES ("done", 1, 1);