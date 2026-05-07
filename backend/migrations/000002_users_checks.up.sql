ALTER TABLE users
ADD CONSTRAINT users_username_format CHECK (
    username ~ '^[a-zA-Z0-9_]{3,32}$'
);


ALTER TABLE users
ADD CONSTRAINT users_email_format CHECK (
    email ~ '^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$'
);


ALTER TABLE users
ADD CONSTRAINT users_email_not_empty CHECK (
    length(trim(email)) > 0
);
