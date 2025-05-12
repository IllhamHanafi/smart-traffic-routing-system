
-- +migrate Up
create table courier
(
    id         uuid primary key,
    name       varchar(255) not null,
    code       varchar(10) unique not null,
    created_at timestamp,
    created_by uuid not null,
    updated_at timestamp,
    updated_by uuid not null
);

-- +migrate Down
drop table courier;