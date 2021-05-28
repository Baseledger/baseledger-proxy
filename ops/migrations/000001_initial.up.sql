DO
$do$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE  rolname = 'baseledger') THEN
      CREATE ROLE baseledger WITH SUPERUSER LOGIN PASSWORD 'ub123';
    END IF;
END
$do$;

SET ROLE baseledger;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

CREATE TABLE public.trustmesh_entries (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    tendermint_block_id text,
    tendermint_transaction_id text,
    tendermint_transaction_timestamp timestamp with time zone,

    sender text,
    receiver text,
    workgroup_id text,

    workstep_type text,
    baseledger_transaction_type text,

    baseledger_transaction_id text,
    referenced_baseledger_transaction_id text,

    business_object_type text,
    baseledger_business_object_id text,
    referenced_baseledger_business_object_id text,

    offchain_process_message_id text,
    referenced_process_message_id text
);

ALTER TABLE public.trustmesh_entries OWNER TO baseledger;