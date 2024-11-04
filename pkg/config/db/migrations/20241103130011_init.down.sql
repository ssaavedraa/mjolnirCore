-- products.down.sql
ALTER TABLE IF EXISTS products DROP CONSTRAINT IF EXISTS products_company_id_fkey;
ALTER TABLE IF EXISTS products DROP CONSTRAINT IF EXISTS products_created_by_fkey;
DROP TABLE IF EXISTS products;

-- users.down.sql
ALTER TABLE IF EXISTS users DROP CONSTRAINT IF EXISTS users_role_id_fkey;
ALTER TABLE IF EXISTS users DROP CONSTRAINT IF EXISTS users_company_id_fkey;
ALTER TABLE IF EXISTS users DROP CONSTRAINT IF EXISTS users_team_id_fkey;
DROP TABLE IF EXISTS users;

-- teams.down.sql
ALTER TABLE IF EXISTS teams DROP CONSTRAINT IF EXISTS teams_company_id_fkey;
DROP TABLE IF EXISTS teams;

-- roles.down.sql
ALTER TABLE IF EXISTS roles DROP CONSTRAINT IF EXISTS roles_company_id_fkey;
DROP TABLE IF EXISTS roles;

-- companies.down.sql
DROP TABLE IF EXISTS companies;
