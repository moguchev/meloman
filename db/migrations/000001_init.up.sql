CREATE TABLE IF NOT EXISTS users (
    id         uuid          NOT NULL,
    login      varchar(64)   NOT NULL,
    password   varchar(256)  NOT NULL,
    salt       int8          NOT NULL,
    created_at timestamp     NOT NULL,
    updated_at timestamp     NOT NULL,
    is_admin   bool          NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS artists (
    id         uuid          NOT NULL,
    full_name  varchar(256)  NOT NULL,
    biography  text          NOT NULL DEFAULT '',
    image      text,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS formats (
    id   uuid        NOT NULL,
    name varchar(64) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS labels (
    id   uuid         NOT NULL,
    name varchar(256) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS albums (
    id        uuid         NOT NULL,
    title     varchar(256) NOT NULL,
    artist_id uuid         NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    year      int2,
    cover     text,
    format_id uuid         REFERENCES formats(id) ON DELETE SET NULL,
    label_id  uuid         REFERENCES labels(id) ON DELETE SET NULL,
    PRIMARY KEY (id),
    UNIQUE (title, artist_id, year, format_id, label_id)
);

CREATE TABLE IF NOT EXISTS tracks (
    id       uuid         NOT NULL,
    number   int2         NOT NULL,
    album_id uuid         NOT NULL REFERENCES albums(id),
    title    varchar(256) NOT NULL,
    length   time         NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users_albums (
    user_id    uuid        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    album_id   uuid        NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    added_at   timestamptz NOT NULL,
    PRIMARY KEY (user_id, album_id)
);
