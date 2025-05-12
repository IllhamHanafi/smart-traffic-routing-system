-- name: GetActiveRoutingLogic :one
SELECT id, allocation_logic FROM routing_decision
WHERE status = 'active' LIMIT 1;

-- name: GetCourierByCode :one
SELECT id, name, code FROM courier
WHERE code = $1;

-- name: InsertRoutingDecisionLog :exec
INSERT INTO routing_decision_log (
    id, order_id, courier_id, routing_decision_id, status, reason, created_at, created_by
) VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7);