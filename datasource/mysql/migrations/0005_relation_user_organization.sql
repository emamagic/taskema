-- +migrate Up
CREATE TABLE user_organization(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    organization_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE user_organization;