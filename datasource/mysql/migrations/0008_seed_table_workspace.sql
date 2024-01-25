-- +migrate Up
INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_one', 1, 1);
INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_two', 1, 1);

INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_one', 2, 1);
INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_two', 2, 1);

INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_one', 3, 1);
