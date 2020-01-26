-- +migrate Up
ALTER TABLE identity ADD COLUMN token varchar(64);
ALTER TABLE identity ADD COLUMN token_expired integer(11);

-- +migrate Down
ALTER TABLE identity DROP COLUMN token;
ALTER TABLE identity DROP COLUMN token_expired;