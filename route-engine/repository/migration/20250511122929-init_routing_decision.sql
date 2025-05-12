
-- +migrate Up
create table routing_decision
(
    id         uuid primary key,
    status     varchar(255) not null,
    allocation_logic json not null,
    created_at timestamp not null,
    created_by uuid not null,
    updated_at timestamp not null,
    updated_by uuid not null
);

-- +migrate Down
drop table routing_decision;