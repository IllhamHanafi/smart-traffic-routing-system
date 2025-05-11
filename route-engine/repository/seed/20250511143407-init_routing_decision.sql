
-- +migrate Up
INSERT INTO routing_decision (id, status, allocation_logic, created_at, created_by, updated_at, updated_by)
VALUES ('2cfc885c-e1ee-4af9-8e35-a9423f1ad8f6', 'active', '{}', '2023-05-11 14:34:07', '11111111-1111-1111-1111-111111111111', '2023-05-11 14:34:07', '11111111-1111-1111-1111-111111111111');

-- +migrate Down
DELETE FROM routing_decision where id = '2cfc885c-e1ee-4af9-8e35-a9423f1ad8f6';