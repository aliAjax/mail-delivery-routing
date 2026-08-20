#!/usr/bin/env sh
set -eu
base="${MAIL_BASE_URL:-http://127.0.0.1:18089}"
curl -fsS "$base/healthz"
curl -fsS "$base/readyz"
curl -fsS "$base/metrics"
curl -fsS -X POST "$base/v1/messages" -H 'Content-Type: application/json' -d '{"tenant_id":"demo","from":"sender@example.test","to":"receiver@example.test","subject":"smoke","text":"hello","idempotency_key":"smoke-1"}'
