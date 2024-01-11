-- +migrate Up
INSERT INTO tasks (`title`, `avatar`, `creator_user_id`, `description`, `board_id`) VALUES ("task_1", "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 1, "description_1", 1);
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`) VALUES ("task_2", 1, "description_1", 1);

INSERT INTO tasks (`title`, `avatar`, `creator_user_id`, `description`, `board_id`) VALUES ("task_1", "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 1, "description_1", 2);
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`) VALUES ("task_2", 1, "description_1", 2);
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`) VALUES ("task_3", 1, "description_1", 2);

INSERT INTO tasks (`title`, `avatar`, `creator_user_id`, `description`, `board_id`) VALUES ("task_1", "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 1, "description_1", 3);