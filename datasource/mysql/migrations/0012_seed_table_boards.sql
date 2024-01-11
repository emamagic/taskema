-- +migrate Up
INSERT INTO boards (`title`, `avatar`, `creator_user_id`, `workspace_id`) VALUES ("board_1", "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 1, 1);
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_2", 1, 1);

INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_1", 1, 2);
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_2", 1, 2);
INSERT INTO boards (`title`, `creator_user_id`, `workspace_id`) VALUES ("board_3", 1, 2);

INSERT INTO boards (`title`, `avatar`, `creator_user_id`, `workspace_id`) VALUES ("board_1", "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 1, 3);