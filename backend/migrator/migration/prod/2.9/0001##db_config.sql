ALTER TABLE db_schema ADD COLUMN IF NOT EXISTS config JSONB NOT NULL DEFAULT '{}';