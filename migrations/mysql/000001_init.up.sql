CREATE TABLE tg_users
(
    id     BIGINT PRIMARY KEY,
    active BOOL DEFAULT true,
    token  VARCHAR(32)
);

CREATE TABLE tokens
(
    token       VARCHAR(32) PRIMARY KEY UNIQUE NOT NULL,
    role        ENUM ('moderator','admin') DEFAULT 'moderator',
    description VARCHAR(255),

    INDEX (token)
);
