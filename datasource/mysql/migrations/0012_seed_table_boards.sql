-- +migrate Up
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_1", 1, 1);
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_2", 1, 1);

INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_1", 1, 2);
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_2", 1, 2);
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_3", 1, 2);

INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_1", 1, 3);