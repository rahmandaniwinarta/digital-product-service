-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS providers;
DROP TABLE IF EXISTS users;

-- +migrate StatementEnd
