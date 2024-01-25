-- +migrate Up
INSERT INTO roles(`title`) VALUES('user');
INSERT INTO roles(`title`) VALUES('board-admin');
INSERT INTO roles(`title`) VALUES('workspace-admin');
INSERT INTO roles(`title`) VALUES('admin');
