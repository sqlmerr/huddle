CREATE TABLE spaces (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 50),
    description VARCHAR(1000) CHECK(char_length(description) BETWEEN 1 AND 1000),
    owner_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_spaces_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);
