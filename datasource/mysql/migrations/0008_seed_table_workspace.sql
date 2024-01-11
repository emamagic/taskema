-- +migrate Up
INSERT INTO workspaces(`title`, `avatar`, `organization_id`, `creator_user_id`) VALUES('workspace_one', "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 1, 1);
INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_two', 1, 1);

INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_one', 2, 1);
INSERT INTO workspaces(`title`, `organization_id`, `creator_user_id`) VALUES('workspace_two', 2, 1);

INSERT INTO workspaces(`title`, `avatar`, `organization_id`, `creator_user_id`) VALUES('workspace_one', "ef026a1a-2e71-4fd6-8e93-66f6e53f79fc", 3, 1);
