
-- +migrate Up
create table routing_decision_log
(
    id                  uuid primary key,
    order_id            uuid not null,
    courier_id          uuid not null,
    routing_decision_id uuid not null,
    status              varchar(255) not null,
    reason              varchar(255),
    created_at          timestamp not null,
    created_by          uuid not null
);

-- +migrate Down
drop table routing_decision_log;