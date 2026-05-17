CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    list_id UUID NOT NULL REFERENCES lists(id) ON DELETE CASCADE,

    title VARCHAR(255) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 255),
    description VARCHAR(1000) CHECK(char_length(description) BETWEEN 1 AND 1000),

    status VARCHAR(50) NOT NULL DEFAULT 'todo' CHECK(char_length(status) BETWEEN 1 AND 50),
    position INT NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)
