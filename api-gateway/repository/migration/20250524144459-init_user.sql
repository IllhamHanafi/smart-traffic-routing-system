
-- +migrate Up
create table user_role
(
    role       varchar(10) primary key
);

create table "user"
(
    id         uuid primary key,
    name       varchar(255) not null,
    role       varchar(10) not null references user_role(role),
    email      varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamp,
    created_by uuid not null,
    updated_at timestamp,
    updated_by uuid not null
);

create unique index user_email_idx on "user"(email);

-- +migrate Down
drop table "user";
drop table user_role;