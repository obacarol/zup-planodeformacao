BEGIN;

ALTER TABLE records DROP COLUMN IF EXISTS transaction_type;

DROP TYPE IF EXISTS transaction_type_enum;

COMMIT;