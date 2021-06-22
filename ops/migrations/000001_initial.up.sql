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

-- TODO: add PK and FK constraints

CREATE TABLE public.organizations (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    organization_name text NOT NULL
);

ALTER TABLE public.organizations OWNER TO baseledger;

CREATE TABLE public.workgroups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    workgroup_name text NOT NULL
);

ALTER TABLE public.workgroups OWNER TO baseledger;

CREATE TABLE public.workgroup_members (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    workgroup_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    organization_endpoint text NOT NULL,
    organization_token text NOT NULL
);

ALTER TABLE public.workgroup_members OWNER TO baseledger;

CREATE TABLE public.trustmesh_entries (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    tendermint_block_id text,
    tendermint_transaction_id text,
    tendermint_transaction_timestamp timestamp with time zone,
    
    type text,

    sender_org_id uuid,
    receiver_org_id uuid,
    workgroup_id uuid,

    workstep_type text,
    baseledger_transaction_type text,

    baseledger_transaction_id text,
    referenced_baseledger_transaction_id text,

    business_object_type text,
    baseledger_business_object_id text,
    referenced_baseledger_business_object_id text,

    offchain_process_message_id text,
    referenced_process_message_id text,

    transaction_status text,
    transaction_hash text
);

ALTER TABLE public.trustmesh_entries OWNER TO baseledger;

CREATE INDEX idx_transaction_status ON public.trustmesh_entries USING btree (transaction_status);