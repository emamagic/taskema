-- +migrate Up
INSERT INTO users(`name`, `email`, `avatar`, `password`) VALUES('test', 'test@gmail.com', "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", '$2a$10$SuPeHxRy3TyIJBqd9hQSceOxs629fL48oYavaiYO45WxOQ.Jz6Xn6');

INSERT INTO organizations(`title`, `creator_user_id`, `avatar`) VALUES('org_one',1, "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc");
INSERT INTO organizations(`title`, `creator_user_id`) VALUES('org_two',1);
INSERT INTO organizations(`title`, `creator_user_id`) VALUES('org_three',1);

INSERT INTO user_organization(`user_id`, `organization_id`) VALUES(1, 1);
INSERT INTO user_organization(`user_id`, `organization_id`) VALUES(1, 2);
INSERT INTO user_organization(`user_id`, `organization_id`) VALUES(1, 3);