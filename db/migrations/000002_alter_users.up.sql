ALTER TABLE users
    DROP COLUMN IF EXISTS is_admin;

CREATE TYPE user_role AS ENUM ('user', 'admin');

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS role user_role NOT NULL DEFAULT 'user';