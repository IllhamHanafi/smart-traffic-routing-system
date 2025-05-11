
-- +migrate Up
create table routing_decision
(
    id         uuid not null,
    status     varchar(255) not null,
    allocation_logic json not null,
    created_at timestamp not null,
    created_by uuid,
    updated_at timestamp not null,
    updated_by uuid
);

-- +migrate Down
drop table routing_decision;