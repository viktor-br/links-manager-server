CREATE TABLE users (
    id         varchar(36),
    username      varchar(255),
    password   varchar(64),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    role       integer,
    PRIMARY KEY(id),
    CONSTRAINT username UNIQUE(username)
);

CREATE TABLE sessions (
    id         varchar(36),
    ip         varchar(45),
    user_id    varchar(36),
    created_at timestamp without time zone,
    expires_at timestamp without time zone,
    PRIMARY KEY(id)
);