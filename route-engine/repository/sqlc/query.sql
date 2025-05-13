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

-- name: InactivateCurrentActiveRoutingDecision :exec
UPDATE routing_decision SET status = 'inactive', updated_at = $1, updated_by = $2
WHERE status = 'active';

-- name: InsertActiveRoutingDecision :exec
INSERT INTO routing_decision (
    id, status, allocation_logic, created_at, created_by, updated_at, updated_by
) VALUES (gen_random_uuid(), 'active', $1, $2, $3, $4, $5);

-- name: GetRoutingDecisionLogs :many
SELECT * FROM routing_decision_log
  WHERE (sqlc.narg(order_id)::uuid is NULL or order_id = sqlc.narg(order_id)::uuid)  
  AND (sqlc.narg(courier_id)::uuid is NULL or courier_id = sqlc.narg(courier_id)::uuid)  
  AND (sqlc.narg(routing_decision_id)::uuid is NULL or routing_decision_id = sqlc.narg(routing_decision_id)::uuid)
  AND (sqlc.narg(status)::text is NULL or status = sqlc.narg(status)::text)
  ORDER BY created_at DESC
  LIMIT $1 OFFSET $2;