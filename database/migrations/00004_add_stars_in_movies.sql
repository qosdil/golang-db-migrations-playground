-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `movies` ADD `stars` VARCHAR(255)  NULL  DEFAULT NULL  AFTER `title`;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `movies` DROP `stars`;