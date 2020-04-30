CREATE TABLE dictionary
(
    id    serial    not null primary key,
    key   varchar   not null unique,
    value varchar   not null,
    time  timestamp not null default now(),
    unique (id, key)
);