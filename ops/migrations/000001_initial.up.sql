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

INSERT INTO public.organizations (organization_name)
VALUES ('Org1'), ('Org2');

CREATE TABLE public.workgroups (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  workgroup_name text NOT NULL
);

ALTER TABLE public.workgroups OWNER TO baseledger;

INSERT INTO public.workgroups (workgroup_name)
VALUES ('Workgroup1');

CREATE TABLE public.workgroup_members (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  workgroup_id uuid NOT NULL,
  organization_id uuid NOT NULL,
  organization_endpoint text NOT NULL,
  organization_token text NOT NULL
);

ALTER TABLE public.workgroup_members OWNER TO baseledger;

INSERT INTO public.workgroup_members (workgroup_id, organization_id, organization_endpoint, organization_token)
WITH
  w AS (
    SELECT id 
    FROM public.workgroups 
  ),
  o AS (
    SELECT id FROM public.organizations
    WHERE organization_name = 'Org1'
  )
  select w.id, o.id, 'host.docker.internal:4222', 'testToken1'
  from w, o;

INSERT INTO public.workgroup_members (workgroup_id, organization_id, organization_endpoint, organization_token)
WITH
  w AS (
    SELECT id 
    FROM public.workgroups 
  ),
  o AS (
    SELECT id FROM public.organizations
    WHERE organization_name = 'Org2'
  )
  select w.id, o.id, 'host.docker.internal:4223', 'testToken1'
  from w, o;

CREATE TABLE public.trustmeshes (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE public.trustmeshes OWNER TO baseledger;
ALTER TABLE ONLY public.trustmeshes ADD CONSTRAINT trustmeshes_pkey PRIMARY KEY (id);

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
  transaction_hash text,
  trustmesh_id uuid NOT NULL
);

ALTER TABLE public.trustmesh_entries OWNER TO baseledger;

CREATE INDEX idx_commitment_state ON public.trustmesh_entries USING btree (commitment_state);

ALTER TABLE ONLY public.trustmesh_entries
  ADD CONSTRAINT trustmesh_entries_trustmesh_id_trustmeshes_id_foreign FOREIGN KEY (trustmesh_id) REFERENCES public.trustmeshes(id) ON UPDATE CASCADE ON DELETE CASCADE;

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

ALTER TABLE ONLY public.offchain_process_messages ADD CONSTRAINT offchain_process_messages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trustmesh_entries ADD CONSTRAINT trustmesh_entries_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trustmesh_entries
  ADD CONSTRAINT trustmesh_entries_offchain_process_message_id_offchain_process_messages_id_foreign FOREIGN KEY (offchain_process_message_id) REFERENCES public.offchain_process_messages(id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE OR REPLACE FUNCTION set_trustmesh_entry_group()
   RETURNS trigger AS
 $$
 DECLARE new_trustmesh_id uuid;
 BEGIN
 	IF NEW.referenced_baseledger_transaction_id = uuid_nil() THEN
 		INSERT INTO trustmeshes VALUES (DEFAULT, DEFAULT) RETURNING id INTO new_trustmesh_id;
	ELSE 
		SELECT trustmesh_id INTO new_trustmesh_id FROM trustmesh_entries WHERE baseledger_transaction_id = NEW.referenced_baseledger_transaction_id;
	END IF;
	NEW.trustmesh_id := new_trustmesh_id;
 RETURN NEW;
 END;
 $$
 LANGUAGE plpgsql;

 CREATE TRIGGER trustmesh_entry_insert_trigger
   BEFORE INSERT
   ON trustmesh_entries
   FOR EACH ROW
   EXECUTE PROCEDURE set_trustmesh_entry_group();
