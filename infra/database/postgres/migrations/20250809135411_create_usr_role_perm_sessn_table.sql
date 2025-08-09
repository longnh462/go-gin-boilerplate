-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE SCHEMA IF NOT EXISTS ggb;
CREATE SCHEMA IF NOT EXISTS keycloak;

SET search_path TO ggb;

-- Roles
CREATE TABLE IF NOT EXISTS roles (
  role_id     uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  role_name   varchar(50)  NOT NULL,
  role_desc   varchar(255),
  is_active   boolean      DEFAULT true,
  create_at   timestamptz  NOT NULL DEFAULT now(),
  update_at   timestamptz  NOT NULL DEFAULT now(),
  create_usr  varchar(25)  NOT NULL DEFAULT 'system',
  update_usr  varchar(25)  NOT NULL DEFAULT 'system',
  deleted_at  timestamptz
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_roles_name ON roles (role_name);
CREATE INDEX IF NOT EXISTS idx_roles_deleted_at ON roles (deleted_at);

-- Permissions
CREATE TABLE IF NOT EXISTS permissions (
  perm_id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  perm_name       varchar(255) NOT NULL,
  perm_desc       varchar(500),
  is_active       boolean      DEFAULT true,
  create_at       timestamptz  NOT NULL DEFAULT now(),
  update_at       timestamptz  NOT NULL DEFAULT now(),
  create_usr      varchar(25)  NOT NULL DEFAULT 'system',
  update_usr      varchar(25)  NOT NULL DEFAULT 'system',
  deleted_at      timestamptz
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_permissions_name ON permissions (perm_name);
CREATE INDEX IF NOT EXISTS idx_permissions_deleted_at ON permissions (deleted_at);

-- Role-Permissions (many-to-many)
CREATE TABLE IF NOT EXISTS role_permissions (
  role_id       uuid NOT NULL,
  perm_id       uuid NOT NULL,
  PRIMARY KEY (role_id, perm_id),
  CONSTRAINT fk_role_permissions_role
    FOREIGN KEY (role_id) REFERENCES ggb.roles(role_id)
    ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT fk_role_permissions_permission
    FOREIGN KEY (perm_id) REFERENCES ggb.permissions(perm_id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

-- Users
CREATE TABLE IF NOT EXISTS users (
  user_id     uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  role_id     uuid         NOT NULL,
  email       varchar(100) NOT NULL,
  username    varchar(50)  NOT NULL,
  password    varchar(100) NOT NULL,
  first_name  varchar(25)  NOT NULL,
  last_name   varchar(50)  NOT NULL,
  phone       varchar(11)  NOT NULL,
  is_active   boolean      DEFAULT false,
  is_verified boolean      DEFAULT false,
  create_at   timestamptz  NOT NULL DEFAULT now(),
  update_at   timestamptz  NOT NULL DEFAULT now(),
  create_usr  varchar(25)  NOT NULL DEFAULT 'system',
  update_usr  varchar(25)  NOT NULL DEFAULT 'system',
  deleted_at  timestamptz,
  CONSTRAINT fk_users_role
    FOREIGN KEY (role_id) REFERENCES ggb.roles(role_id)
    ON UPDATE CASCADE ON DELETE RESTRICT
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_users_email ON users (email);
CREATE UNIQUE INDEX IF NOT EXISTS ux_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_role_id ON users (role_id);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

-- Sessions
CREATE TABLE IF NOT EXISTS sessions (
  session_id  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id     uuid         NOT NULL,
  token       varchar(255) NOT NULL,
  expires_at  timestamptz  NOT NULL,
  last_active timestamptz  NOT NULL,
  ip_address  varchar(45)  NOT NULL,
  user_agent  varchar(255) NOT NULL,
  is_active   boolean      DEFAULT true,
  create_at   timestamptz  NOT NULL DEFAULT now(),
  update_at   timestamptz  NOT NULL DEFAULT now(),
  create_usr  varchar(25)  NOT NULL DEFAULT 'system',
  update_usr  varchar(25)  NOT NULL DEFAULT 'system',
  deleted_at  timestamptz,
  CONSTRAINT fk_sessions_user
    FOREIGN KEY (user_id) REFERENCES ggb.users(user_id)
    ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions (user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_deleted_at ON sessions (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ggb.sessions;
DROP TABLE IF EXISTS ggb.users;
DROP TABLE IF EXISTS ggb.role_permissions;
DROP TABLE IF EXISTS ggb.permissions;
DROP TABLE IF EXISTS ggb.roles;
-- +goose StatementEnd