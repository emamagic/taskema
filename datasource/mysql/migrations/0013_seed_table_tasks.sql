-- +migrate Up
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`, `assigned_user_id`) VALUES ("task_1", 1, "description_1", 1, 1);
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`, `assigned_user_id`) VALUES ("task_2", 1, "description_1", 1, 1);

INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`, `assigned_user_id`) VALUES ("task_1", 1, "description_1", 2, 1);
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`, `assigned_user_id`) VALUES ("task_2", 1, "description_1", 2, 1);
INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`, `assigned_user_id`) VALUES ("task_3", 1, "description_1", 2, 1);

INSERT INTO tasks (`title`, `creator_user_id`, `description`, `board_id`, `assigned_user_id`) VALUES ("task_1", 1, "description_1", 3, 1);