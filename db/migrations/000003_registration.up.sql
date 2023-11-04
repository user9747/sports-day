CREATE TABLE IF NOT EXISTS registrations(
    user_id UUID REFERENCES users(id),
    event_id INT REFERENCES events(id),
    created_at timestamp default now(),
    PRIMARY KEY (user_id, event_id)
)