-- +migrate Up
CREATE TABLE roles (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE roles;