-- +migrate Up
INSERT INTO users(`name`, `email`, `password`) VALUES('test', 'test@gmail.com', '$2a$10$SuPeHxRy3TyIJBqd9hQSceOxs629fL48oYavaiYO45WxOQ.Jz6Xn6');

INSERT INTO organizations(`title`, `creator_user_id`) VALUES('org_one',1);
INSERT INTO organizations(`title`, `creator_user_id`) VALUES('org_two',1);
INSERT INTO organizations(`title`, `creator_user_id`) VALUES('org_three',1);

INSERT INTO user_organization(`user_id`, `organization_id`) VALUES(1, 1);
INSERT INTO user_organization(`user_id`, `organization_id`) VALUES(1, 2);
INSERT INTO user_organization(`user_id`, `organization_id`) VALUES(1, 3);