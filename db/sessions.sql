CREATE TABLE sessions (
    id         varchar(36),
    ip         varchar(45),
    user_id    varchar(36),
    created_at timestamp without time zone,
    expires_at timestamp without time zone,
    PRIMARY KEY(id)
);