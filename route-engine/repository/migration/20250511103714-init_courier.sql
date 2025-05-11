
-- +migrate Up
create table courier
(
    id         uuid primary key,
    name       varchar(255),
    code       varchar(10),
    created_at timestamp,
    created_by uuid,
    updated_at timestamp,
    updated_by uuid
);

-- +migrate Down
drop table courier;