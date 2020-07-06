
-- +migrate Up
    ALTER TABLE identity MODIFY COLUMN create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +migrate Down
    ALTER TABLE identity MODIFY COLUMN create_date DATETIME NOT NULL;