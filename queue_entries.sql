CREATE TABLE queue_entries (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticket_id   TEXT UNIQUE NOT NULL,
    user_id     TEXT NOT NULL,
    queue_id    TEXT NOT NULL,
    position    INT NOT NULL,
    status      TEXT DEFAULT 'waiting',
    created_at  TIMESTAMP DEFAULT NOW()
);
