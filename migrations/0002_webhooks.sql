CREATE TABLE webhook_deliveries (id text PRIMARY KEY, message_id text NOT NULL, url text NOT NULL, attempts integer NOT NULL DEFAULT 0, status text NOT NULL, next_attempt timestamptz);
CREATE INDEX messages_tenant_status_idx ON messages(tenant_id, status);
