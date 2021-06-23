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
    tendermint_transaction_id uuid,
    tendermint_transaction_timestamp timestamp with time zone,
    
    entry_type text,

    sender_org_id uuid,
    receiver_org_id uuid,
    workgroup_id uuid,

    workstep_type text,
    baseledger_transaction_type text,

    baseledger_transaction_id uuid,
    referenced_baseledger_transaction_id uuid,

    business_object_type text,
    baseledger_business_object_id uuid,
    referenced_baseledger_business_object_id uuid,

    offchain_process_message_id uuid,
    referenced_process_message_id uuid,

    commitment_state text,
    transaction_hash text
);

ALTER TABLE public.trustmesh_entries OWNER TO baseledger;

CREATE INDEX idx_commitment_state ON public.trustmesh_entries USING btree (commitment_state);

CREATE TABLE public.offchain_process_messages (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    sender_id uuid,
	receiver_id uuid,
	topic text,
	referenced_offchain_process_message_id uuid,
	baseledger_sync_tree_json text,
	workstep_type text,
	business_object_proof text,
	tendermint_transaction_id_of_stored_proof uuid,
	baseledger_transaction_id_of_stored_proof uuid,
    baseledger_business_object_id uuid,
	referenced_baseledger_business_object_id uuid,
	status_text_message text,
    business_object_type text,
	baseledger_transaction_type text,
	referenced_baseledger_transaction_id uuid,
	entry_type text
);

ALTER TABLE public.offchain_process_messages OWNER TO baseledger;

ALTER TABLE ONLY public.offchain_process_messages ADD CONSTRAINT offchain_process_messagess_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trustmesh_entries ADD CONSTRAINT trustmesh_entries_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trustmesh_entries
    ALTER COLUMN offchain_process_message_id TYPE uuid USING (uuid_generate_v4());

ALTER TABLE ONLY public.trustmesh_entries
    ADD CONSTRAINT trustmesh_entries_offchain_process_message_id_offchain_process_messages_id_foreign FOREIGN KEY (offchain_process_message_id) REFERENCES public.offchain_process_messages(id) ON UPDATE CASCADE ON DELETE CASCADE;