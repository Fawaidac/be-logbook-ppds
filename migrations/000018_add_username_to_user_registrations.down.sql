-- Rollback: drop username column from user_registrations
ALTER TABLE user_registrations DROP COLUMN IF EXISTS username;
