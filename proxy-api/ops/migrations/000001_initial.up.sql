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
        SELECT trustmesh_id INTO new_trustmesh_id FROM trustmesh_entries WHERE baseledger_transaction_id = NEW.referenced_baseledger_transaction_id LIMIT 1;
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


-- Add trustmeshes for testing purposes

INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('393bce3e-d014-4dde-afac-26048cef3afc', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 1}', 'Initial', '4a605ebd0a750d923203675ffc4e47e5', '247edc6d-c087-4717-bd47-6cbaa2275305', '247edc6d-c087-4717-bd47-6cbaa2275305', 'a87d4a07-dfa5-4f0b-9f39-24a06b47cdd5', '00000000-0000-0000-0000-000000000000', 'Initial suggested', 'PurchaseOrder', 'Suggest', '00000000-0000-0000-0000-000000000000', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('0ca0f164-20fc-4e3f-8fd4-aacc6cd9969b', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 1}', 'Feedback', '', '150da2c6-89ff-4630-be27-cf3e77d9bfd5', '150da2c6-89ff-4630-be27-cf3e77d9bfd5', '00000000-0000-0000-0000-000000000000', 'a87d4a07-dfa5-4f0b-9f39-24a06b47cdd5', '', 'PurchaseOrder', 'Reject', '247edc6d-c087-4717-bd47-6cbaa2275305', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('3730a017-ed2c-4de5-92df-8c37a7c62212', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 2}', 'NewVersion', 'aea10bb921e62cac171c4a49b0c60d85', 'addfce8d-7e31-4016-ae00-b287732b88c4', 'addfce8d-7e31-4016-ae00-b287732b88c4', '3e6293fc-6aca-45ba-8537-9a81594c8736', 'd9818659-dd80-4aa7-9c76-3a6e7fbb75b1', 'NewVersion suggested', 'PurchaseOrder', 'Suggest', '150da2c6-89ff-4630-be27-cf3e77d9bfd5', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('7cf6f3a2-56ff-43d3-8dee-ea0c09a4d608', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 1}', 'Feedback', '', '856d96aa-f0d1-4548-a815-54ea39c880ca', '856d96aa-f0d1-4548-a815-54ea39c880ca', '00000000-0000-0000-0000-000000000000', '3e6293fc-6aca-45ba-8537-9a81594c8736', '', 'PurchaseOrder', 'Approve', 'addfce8d-7e31-4016-ae00-b287732b88c4', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('ac2fcb39-32bf-4595-8cd8-a3beae674651', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 3}', 'NextWorkstep', '04793029d736560b66d0b7df9332c3af', '2f0c188f-74ea-47aa-8b93-9fa79af5e4df', '2f0c188f-74ea-47aa-8b93-9fa79af5e4df', '20a8b57c-535e-4137-91fe-d94584758c7f', 'db714dec-4544-40f8-81cf-9de00e20f336', 'NextWorkstep suggested', 'PurchaseOrder', 'Suggest', '856d96aa-f0d1-4548-a815-54ea39c880ca', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('d56c8c22-8004-46e2-aa2f-846b3660ad8d', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 1}', 'Feedback', '', '537744b1-50db-4d32-a14a-5504568e74c2', '537744b1-50db-4d32-a14a-5504568e74c2', '00000000-0000-0000-0000-000000000000', '20a8b57c-535e-4137-91fe-d94584758c7f', '', 'PurchaseOrder', 'Reject', '2f0c188f-74ea-47aa-8b93-9fa79af5e4df', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('935d1bd6-5143-4b56-b949-3c3fdfdb46e0', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 4}', 'NewVersion', '1da8510c1c06bcdefec063c2e719b64f', 'e8c9dfdd-524a-418a-bb36-acc4e77f7aaa', 'e8c9dfdd-524a-418a-bb36-acc4e77f7aaa', '5d8a4704-cd53-4b6a-bcea-dc58b05cf630', 'd4d225ec-b26d-4aa9-8889-3890b26e34f0', 'NewVersion suggested', 'PurchaseOrder', 'Suggest', '537744b1-50db-4d32-a14a-5504568e74c2', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('325753da-6bbb-4853-b9fb-09ef00fc907c', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 4}', 'Feedback', '', '414f19fa-20cf-4818-9089-9a6d0e01bd3b', '414f19fa-20cf-4818-9089-9a6d0e01bd3b', '00000000-0000-0000-0000-000000000000', '5d8a4704-cd53-4b6a-bcea-dc58b05cf630', '', 'PurchaseOrder', 'Approve', '537744b1-50db-4d32-a14a-5504568e74c2', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('532c6d65-3b75-4908-8f0c-eca1c8bcccc6', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 5}', 'Final', 'ff07a85551d2cf10289dc672b9c1da81', '9eb0e15e-d38b-474c-a571-509026abd766', '9eb0e15e-d38b-474c-a571-509026abd766', '169f104f-980e-42bb-a128-73daf259bc39', '09d2f10a-dce8-4022-94d8-98c1010f8f60', 'FinalWorkstep suggested', 'PurchaseOrder', 'Suggest', '414f19fa-20cf-4818-9089-9a6d0e01bd3b', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('0ccc0e6c-24e9-4308-9e2e-bff69858c539', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{PurchaseOrderID: 5}', 'Feedback', '', '85355693-e70d-43ef-bec0-33ad2e39ee87', '85355693-e70d-43ef-bec0-33ad2e39ee87', '00000000-0000-0000-0000-000000000000', '169f104f-980e-42bb-a128-73daf259bc39', '', 'PurchaseOrder', 'Approve', '00000000-0000-0000-0000-000000000000', 'FeedbackSent');


INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('23300b78-e5ed-4ea2-9da9-0f52b04b4b2e', '11', '247edc6d-c087-4717-bd47-6cbaa2275305', '2021-07-12 21:13:15.102635+00', 'SuggestionSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Initial', 'Suggest', '247edc6d-c087-4717-bd47-6cbaa2275305', '00000000-0000-0000-0000-000000000000', 'PurchaseOrder', 'a87d4a07-dfa5-4f0b-9f39-24a06b47cdd5', '00000000-0000-0000-0000-000000000000', '393bce3e-d014-4dde-afac-26048cef3afc', '00000000-0000-0000-0000-000000000000', 'COMMITTED', 'EB878241F26B916F809AB4E707063F1E219568CBEBDE02D0B18FFCDDE18B4366', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('9c21b541-b609-486d-bbc1-fc4d649bb5b0', '41', '150da2c6-89ff-4630-be27-cf3e77d9bfd5', '2021-07-12 21:18:15.649712+00', 'FeedbackSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Reject', '150da2c6-89ff-4630-be27-cf3e77d9bfd5', '247edc6d-c087-4717-bd47-6cbaa2275305', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', 'a87d4a07-dfa5-4f0b-9f39-24a06b47cdd5', '0ca0f164-20fc-4e3f-8fd4-aacc6cd9969b', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '16203E04729D18E158A72A21D792D35D2345F7CD347EAF18B0E409CC923C7E7F', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('374354c0-cc2b-49ac-8018-f9c281b1d47a', '65', 'addfce8d-7e31-4016-ae00-b287732b88c4', '2021-07-12 21:22:16.082541+00', 'SuggestionSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'NewVersion', 'Suggest', 'addfce8d-7e31-4016-ae00-b287732b88c4', '150da2c6-89ff-4630-be27-cf3e77d9bfd5', 'PurchaseOrder', '3e6293fc-6aca-45ba-8537-9a81594c8736', 'd9818659-dd80-4aa7-9c76-3a6e7fbb75b1', '3730a017-ed2c-4de5-92df-8c37a7c62212', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '8F76BE32186108A1254530FD3E0059125E033971178C8D6C46E7AC1E4E863835', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('814a46b7-5890-49eb-b7e9-156b9f6e302c', '72', '856d96aa-f0d1-4548-a815-54ea39c880ca', '2021-07-12 21:23:26.211709+00', 'FeedbackSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Approve', '856d96aa-f0d1-4548-a815-54ea39c880ca', 'addfce8d-7e31-4016-ae00-b287732b88c4', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', '3e6293fc-6aca-45ba-8537-9a81594c8736', '7cf6f3a2-56ff-43d3-8dee-ea0c09a4d608', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '08BBD662045F31A7B16028776672DC19FBBF2FBBE34737EE30B5D97D862848BF', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('71183739-2d1c-452c-92ad-2b86894725dc', '79', '2f0c188f-74ea-47aa-8b93-9fa79af5e4df', '2021-07-12 21:24:36.337085+00', 'SuggestionSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'NextWorkstep', 'Suggest', '2f0c188f-74ea-47aa-8b93-9fa79af5e4df', '856d96aa-f0d1-4548-a815-54ea39c880ca', 'PurchaseOrder', '20a8b57c-535e-4137-91fe-d94584758c7f', 'db714dec-4544-40f8-81cf-9de00e20f336', 'ac2fcb39-32bf-4595-8cd8-a3beae674651', '00000000-0000-0000-0000-000000000000', 'COMMITTED', 'B94C0F900469B6473A5F51A4255348F0E5E73BEDB8571CA60A08E216F4E3B319', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('18f942af-0922-4b73-8865-b664beb08088', '84', '537744b1-50db-4d32-a14a-5504568e74c2', '2021-07-12 21:25:26.425783+00', 'FeedbackSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Reject', '537744b1-50db-4d32-a14a-5504568e74c2', '2f0c188f-74ea-47aa-8b93-9fa79af5e4df', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', '20a8b57c-535e-4137-91fe-d94584758c7f', 'd56c8c22-8004-46e2-aa2f-846b3660ad8d', '00000000-0000-0000-0000-000000000000', 'COMMITTED', 'CD42A75C128990AA548CA9097C388551022C4AB0DC2430A2D50E2F4DB845DB4C', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('ff97a2a8-598f-4c8a-9d4a-7b95d700d173', '90', 'e8c9dfdd-524a-418a-bb36-acc4e77f7aaa', '2021-07-12 21:26:26.537006+00', 'SuggestionSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'NewVersion', 'Suggest', 'e8c9dfdd-524a-418a-bb36-acc4e77f7aaa', '537744b1-50db-4d32-a14a-5504568e74c2', 'PurchaseOrder', '5d8a4704-cd53-4b6a-bcea-dc58b05cf630', 'd4d225ec-b26d-4aa9-8889-3890b26e34f0', '935d1bd6-5143-4b56-b949-3c3fdfdb46e0', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '5BF8F866A0A41300F9EC227B6B45E7FEEB7DD7E909DF61C67FD46AE82AD60571', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('c58ea08a-e7ba-48f6-bb5b-2b87aefd82e6', '95', '414f19fa-20cf-4818-9089-9a6d0e01bd3b', '2021-07-12 21:27:16.622057+00', 'FeedbackSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Approve', '414f19fa-20cf-4818-9089-9a6d0e01bd3b', 'e8c9dfdd-524a-418a-bb36-acc4e77f7aaa', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', '5d8a4704-cd53-4b6a-bcea-dc58b05cf630', '325753da-6bbb-4853-b9fb-09ef00fc907c', '00000000-0000-0000-0000-000000000000', 'COMMITTED', 'A388DD0B120A26C7A46BCFD4E4CC1DD2C8136073B7E275FB384F77369A7CF77A', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('9b78a73c-e642-4b77-8125-39e05ecb1b4a', '108', '9eb0e15e-d38b-474c-a571-509026abd766', '2021-07-12 21:29:26.855506+00', 'SuggestionSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Final', 'Suggest', '9eb0e15e-d38b-474c-a571-509026abd766', '414f19fa-20cf-4818-9089-9a6d0e01bd3b', 'PurchaseOrder', '169f104f-980e-42bb-a128-73daf259bc39', '09d2f10a-dce8-4022-94d8-98c1010f8f60', '532c6d65-3b75-4908-8f0c-eca1c8bcccc6', '00000000-0000-0000-0000-000000000000', 'COMMITTED', 'EA26D958612910637467BC24162F05738DC6B195DB4D4E10C8ACC00CB965B37C', '189d9a8c-a1f9-4069-bbe6-905ac644600b');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('e56df304-aef8-4f45-a4b4-fa8ae1fa98c2', '112', '85355693-e70d-43ef-bec0-33ad2e39ee87', '2021-07-12 21:30:06.93092+00', 'FeedbackSent', '5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Approve', '85355693-e70d-43ef-bec0-33ad2e39ee87', '9eb0e15e-d38b-474c-a571-509026abd766', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', '169f104f-980e-42bb-a128-73daf259bc39', '0ccc0e6c-24e9-4308-9e2e-bff69858c539', '00000000-0000-0000-0000-000000000000', 'COMMITTED', 'D3689177A78A6C4ADE8008A0E9D2F2E0C4B9CDF973C9132E32D89BF9C0567516', '189d9a8c-a1f9-4069-bbe6-905ac644600b');