-- Create tables with initial structure, including deleted_at column for soft deletes
CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    domain VARCHAR(255),
    nit VARCHAR(255) UNIQUE,
    address VARCHAR(255),
    phone_number VARCHAR(255),
    is_draft BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP  -- Add deleted_at column for soft delete
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    image_url VARCHAR(255),
    created_by INTEGER NOT NULL,
    company_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP  -- Add deleted_at column for soft delete
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    company_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP  -- Add deleted_at column for soft delete
);

CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    company_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP  -- Add deleted_at column for soft delete
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(255),
    address VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    is_draft BOOLEAN DEFAULT FALSE,
    invite_id VARCHAR(255) DEFAULT NULL,
    role_id INTEGER,
    company_id INTEGER,
    team_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP  -- Add deleted_at column for soft delete
);

-- Add conditional foreign key constraints
DO $$
BEGIN
    -- Products table foreign key to companies table
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_products_company'
    ) THEN
        ALTER TABLE products
        ADD CONSTRAINT fk_products_company
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;
    END IF;

    -- Products table foreign key to users table for created_by
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_products_created_by'
    ) THEN
        ALTER TABLE products
        ADD CONSTRAINT fk_products_created_by
        FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE;
    END IF;

    -- Roles table foreign key to companies table
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_roles_company'
    ) THEN
        ALTER TABLE roles
        ADD CONSTRAINT fk_roles_company
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;
    END IF;

    -- Teams table foreign key to companies table
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_teams_company'
    ) THEN
        ALTER TABLE teams
        ADD CONSTRAINT fk_teams_company
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL;
    END IF;

    -- Users table foreign key to roles table
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_users_role'
    ) THEN
        ALTER TABLE users
        ADD CONSTRAINT fk_users_role
        FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL;
    END IF;

    -- Users table foreign key to companies table
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_users_company'
    ) THEN
        ALTER TABLE users
        ADD CONSTRAINT fk_users_company
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL;
    END IF;

    -- Users table foreign key to teams table
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'fk_users_team'
    ) THEN
        ALTER TABLE users
        ADD CONSTRAINT fk_users_team
        FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE SET NULL;
    END IF;
END $$;
