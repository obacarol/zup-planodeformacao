BEGIN;

ALTER TABLE account DROP COLUMN IF EXISTS name;

ALTER TABLE account DROP COLUMN IF EXISTS cpf;

ALTER TABLE account DROP COLUMN IF EXISTS balance_account;

COMMIT;