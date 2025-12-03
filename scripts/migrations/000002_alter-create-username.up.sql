ALTER TABLE users
ADD COLUMN username VARCHAR(250);

ALTER TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);
