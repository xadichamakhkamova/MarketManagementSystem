-- name: CreateDebt :one
INSERT INTO debts 
    (
        first_name, 
        last_name, 
        phone_number, 
        jshshir, 
        address, 
        bag_id,
        price,
        price_paid,
        acquaintance, 
        collateral, 
        deadline
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING
    id,
    first_name, 
    last_name, 
    phone_number, 
    jshshir, 
    address, 
    bag_id,
    price,
    price_paid,
    acquaintance, 
    collateral, 
    deadline,
    created_at,
    updated_at,
    deleted_at;


-- name: UpdateDebt :one
UPDATE debts
SET 
    first_name = COALESCE($2, first_name),
    last_name = COALESCE($3, last_name),
    phone_number = COALESCE($4, phone_number),
    jshshir = COALESCE($5, jshshir),
    address = COALESCE($6, address),
    bag_id = COALESCE($7, bag_id),
    price = COALESCE($8, price),
    price_paid = COALESCE($9, price_paid),
    acquaintance = COALESCE($10, acquaintance),
    collateral = COALESCE($11, collateral),
    deadline = COALESCE($12, deadline),
    updated_at = $13
WHERE id = $1
RETURNING
    id,
    first_name,
    last_name,
    phone_number,
    jshshir,
    address,
    bag_id,
    price,
    price_paid,
    acquaintance,
    collateral,
    deadline,
    created_at,
    updated_at,
    deleted_at;
    
-- name: DeleteDebt :exec
UPDATE debts
SET deleted_at = $2
WHERE id = $1;

-- name: GetDebtById :one
SELECT * FROM debts
WHERE id = $1;


-- name: GetDebtByFilter :many
SELECT * FROM debts
WHERE 
    ($1::text = 'allDataFromDB' OR 
    first_name ILIKE '%' || $1 || '%' OR
    last_name ILIKE '%' || $1 || '%' OR
    phone_number ILIKE '%' || $1 || '%' OR
    jshshir ILIKE '%' || $1 || '%' OR
    address ILIKE '%' || $1 || '%' OR
    acquaintance ILIKE '%' || $1 || '%' OR
    collateral ILIKE '%' || $1 || '%' OR
    CAST(deadline AS text) ILIKE '%' || $1 || '%');


