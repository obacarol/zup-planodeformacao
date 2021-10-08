BEGIN;

CREATE TYPE transaction_type_enum AS ENUM (
    'Pay in',
    'Withdrawal',
    'Transfer'
    );

ALTER TABLE records ADD COLUMN transaction_type transaction_type_enum;

COMMIT;