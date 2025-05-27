
-- +migrate Up
insert into user_role(role)
values ('ADMIN'),
       ('USER');

-- +migrate Down
delete from user_role where role in ('ADMIN', 'USER');