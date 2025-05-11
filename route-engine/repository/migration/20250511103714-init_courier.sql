
-- +migrate Up
create table courier
(
    id         bigint primary key,
    name       varchar(255),
    created_at timestamp,
    updated_at timestamp
);

-- +migrate Down
drop table courier;