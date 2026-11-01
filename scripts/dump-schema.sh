#!/bin/bash

CONTAINER="my-postgres"
DB_USER="cooldocs"
DB_NAME="cooldocs"
# OUTPUT="schema.sql"

echo "Generating PostgreSQL schema dump..."

docker exec "$CONTAINER" \
  pg_dump -U "$DB_USER" -d "$DB_NAME" \
  --schema-only \
  --no-owner \
  --no-privileges

if [ $? -eq 0 ]; then
  echo "Schema dump successfully generated: $OUTPUT"
else
  echo "Schema dump failed!"
  exit 1
fi
