-- name: UpsertProductSales :one
INSERT INTO dashboard (product_id, color, cost_price, selling_price) 
VALUES ($1, $2, $3, $4) 
RETURNING 'success' AS result;

-- name: GetDashboardReport :one
SELECT
    COALESCE(SUM(selling_price * units_sold), 0) :: BIGINT AS total_sales,
    COALESCE(SUM(cost_price * units_sold), 0) :: BIGINT AS total_cost_price,
    COALESCE(SUM((selling_price - cost_price) * units_sold), 0) :: BIGINT AS total_net_profit,
    COALESCE(SUM(units_sold), 0) :: BIGINT AS total_units_sold
FROM dashboard;
