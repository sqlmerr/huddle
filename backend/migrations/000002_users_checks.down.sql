-- Remove CHECK constraints in reverse order
ALTER TABLE users
DROP CONSTRAINT IF EXISTS users_email_not_empty;

ALTER TABLE users
DROP CONSTRAINT IF EXISTS users_email_format;

ALTER TABLE users
DROP CONSTRAINT IF EXISTS users_username_format;
