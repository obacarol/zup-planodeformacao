CREATE SEQUENCE account_id_seq1;

CREATE SEQUENCE records_id_seq1;

CREATE TABLE public.account
(
    id integer NOT NULL DEFAULT nextval('account_id_seq1'::regclass),
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    cpf bigint NOT NULL,
    creation_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    balance_account numeric NOT NULL DEFAULT 0,
    CONSTRAINT account_pkey1 PRIMARY KEY (id)
);

CREATE TABLE public.records
(
    id integer NOT NULL DEFAULT nextval('records_id_seq1'::regclass),
    id_account_from integer NOT NULL,
	id_account_to integer,
    transaction_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    transaction_value numeric,
    CONSTRAINT record_pkey1 PRIMARY KEY (id),
    CONSTRAINT records_id_account_from_fkey FOREIGN KEY (id_account_from)
        REFERENCES public.account (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
	CONSTRAINT records_id_account_to_fkey FOREIGN KEY (id_account_to)
	REFERENCES public.account (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);