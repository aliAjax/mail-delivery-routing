CREATE TABLE tenants (id text PRIMARY KEY, name text NOT NULL, daily_quota integer NOT NULL DEFAULT 0, paused boolean NOT NULL DEFAULT false);
CREATE TABLE messages (id text PRIMARY KEY, tenant_id text NOT NULL, sender text NOT NULL, recipient text NOT NULL, subject text NOT NULL, status text NOT NULL, created_at timestamptz NOT NULL);
CREATE TABLE delivery_attempts (id bigserial PRIMARY KEY, message_id text NOT NULL, attempt integer NOT NULL, status text NOT NULL, reason text, created_at timestamptz NOT NULL);
