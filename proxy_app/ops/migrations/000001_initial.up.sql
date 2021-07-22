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

INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('d61a4a37-6ab6-40da-a9c2-5af3024014ca', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{"RootProof":"afca87a34a5318583e45daf14fce7599","Nodes":[{"SyncTreeNodeID":"","Value":"MerkleTreeForExiting:Inv123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"afca87a34a5318583e45daf14fce7599","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":1,"Index":0}]}', 'FinalWorkstep', 'afca87a34a5318583e45daf14fce7599', 'd130279e-6373-4744-82f9-fc542183eadc', 'd130279e-6373-4744-82f9-fc542183eadc', '12794521-4d92-40f6-b28f-c7baa282a7c7', '94dd403e-e1e6-4834-826a-5a87a471b44e', 'FinalWorkstep suggested', 'MerkleTreeForExiting', 'Suggest', '01f9a344-9937-410e-b415-455105c4567b', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('118118dd-ff64-439f-b0bf-f02b443654d2', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'd61a4a37-6ab6-40da-a9c2-5af3024014ca', '{"RootProof":"afca87a34a5318583e45daf14fce7599","Nodes":[{"SyncTreeNodeID":"","Value":"MerkleTreeForExiting:Inv123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"afca87a34a5318583e45daf14fce7599","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":1,"Index":0}]}', 'Feedback', '', '83b858a8-8cbc-4f54-ac4b-422b5c9ccc5e', '83b858a8-8cbc-4f54-ac4b-422b5c9ccc5e', '00000000-0000-0000-0000-000000000000', '12794521-4d92-40f6-b28f-c7baa282a7c7', '', 'MerkleTreeForExiting', 'Approve', 'd130279e-6373-4744-82f9-fc542183eadc', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('6a55d674-3ac9-47a3-a83e-b354c623936f', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{"RootProof":"670a71cd92e66848a22048424699bc33","Nodes":[{"SyncTreeNodeID":"","Value":"PurchaseOrderID:PO123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"Currency:EUR","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"MaterialID:4711","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"Quantity:3","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"411b58f44e9444b6fee2d931f209a01b","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"3fb82c66ce419979ec3612d35b69bf36","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"670a71cd92e66848a22048424699bc33","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'Initial', '670a71cd92e66848a22048424699bc33', '467dca32-ba2a-4352-9685-6144df7336e4', '467dca32-ba2a-4352-9685-6144df7336e4', 'b1ade3ef-76aa-45c0-9c08-8c9c295f3601', '00000000-0000-0000-0000-000000000000', 'Initial suggested', 'PurchaseOrder', 'Suggest', '00000000-0000-0000-0000-000000000000', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('89f90a49-78c2-4dc9-952a-1c130d5e75f5', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '6a55d674-3ac9-47a3-a83e-b354c623936f', '{"RootProof":"670a71cd92e66848a22048424699bc33","Nodes":[{"SyncTreeNodeID":"","Value":"PurchaseOrderID:PO123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"Currency:EUR","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"MaterialID:4711","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"Quantity:3","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"411b58f44e9444b6fee2d931f209a01b","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"3fb82c66ce419979ec3612d35b69bf36","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"670a71cd92e66848a22048424699bc33","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'Feedback', '', 'c0f5591b-fbf7-483b-bf1e-c18e362c7a06', 'c0f5591b-fbf7-483b-bf1e-c18e362c7a06', '00000000-0000-0000-0000-000000000000', 'b1ade3ef-76aa-45c0-9c08-8c9c295f3601', '', 'PurchaseOrder', 'Reject', '467dca32-ba2a-4352-9685-6144df7336e4', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('e3584552-7582-4b21-814a-bd8fe0fe7e39', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{"RootProof":"6ee5c3b0e6776169247fa0bec3fcce00","Nodes":[{"SyncTreeNodeID":"","Value":"Currency:USD","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"MaterialID:4711","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"Quantity:5","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"PurchaseOrderID:PO123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"10a75c66085336059930b5155d57e1eb","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"aca524896a245c336a78e0ea6a3ba125","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"6ee5c3b0e6776169247fa0bec3fcce00","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'NewVersion', '6ee5c3b0e6776169247fa0bec3fcce00', '7fc8d994-0c3a-4cea-9946-828d5d60a99e', '7fc8d994-0c3a-4cea-9946-828d5d60a99e', 'beda5b1f-0e6d-4cd6-b960-f9c0229e1d47', 'b1ade3ef-76aa-45c0-9c08-8c9c295f3601', 'NewVersion suggested', 'PurchaseOrder', 'Suggest', 'c0f5591b-fbf7-483b-bf1e-c18e362c7a06', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('6a03eca3-e92c-48c8-a55b-2f3a3794495c', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'e3584552-7582-4b21-814a-bd8fe0fe7e39', '{"RootProof":"6ee5c3b0e6776169247fa0bec3fcce00","Nodes":[{"SyncTreeNodeID":"","Value":"Currency:USD","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"MaterialID:4711","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"Quantity:5","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"PurchaseOrderID:PO123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"10a75c66085336059930b5155d57e1eb","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"aca524896a245c336a78e0ea6a3ba125","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"6ee5c3b0e6776169247fa0bec3fcce00","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'Feedback', '', 'ce7f46b0-a1bc-48be-8b4a-95affab4a89b', 'ce7f46b0-a1bc-48be-8b4a-95affab4a89b', '00000000-0000-0000-0000-000000000000', 'beda5b1f-0e6d-4cd6-b960-f9c0229e1d47', '', 'PurchaseOrder', 'Approve', '7fc8d994-0c3a-4cea-9946-828d5d60a99e', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('05aa6e2b-5af4-49ba-9fac-07ab442023a8', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{"RootProof":"f7f759688b10a164a3b2c314dad0d496","Nodes":[{"SyncTreeNodeID":"","Value":"InvoiceID:Inv123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"Currency:USD","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"Amount:5000","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"7f1b22cdb9ccda7de6a3bb94d21e3189","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"ae05f4ec3e54d400d186e07947932d95","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"f7f759688b10a164a3b2c314dad0d496","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'NextWorkstep', 'f7f759688b10a164a3b2c314dad0d496', 'e4a12657-a23b-48b1-98a8-06f7bdaa0f7f', 'e4a12657-a23b-48b1-98a8-06f7bdaa0f7f', 'ef9ea07d-6294-4a13-9ac1-9783190fa794', 'beda5b1f-0e6d-4cd6-b960-f9c0229e1d47', 'NextWorkstep suggested', 'Invoice', 'Suggest', 'ce7f46b0-a1bc-48be-8b4a-95affab4a89b', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('52c1bfb1-bc1b-417a-a204-160cc02a56ec', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '05aa6e2b-5af4-49ba-9fac-07ab442023a8', '{"RootProof":"f7f759688b10a164a3b2c314dad0d496","Nodes":[{"SyncTreeNodeID":"","Value":"InvoiceID:Inv123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"Currency:USD","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"Amount:5000","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"7f1b22cdb9ccda7de6a3bb94d21e3189","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"ae05f4ec3e54d400d186e07947932d95","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"f7f759688b10a164a3b2c314dad0d496","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'Feedback', '', '5c4f4f57-ac12-4ae8-a8d3-f290870f1aee', '5c4f4f57-ac12-4ae8-a8d3-f290870f1aee', '00000000-0000-0000-0000-000000000000', 'ef9ea07d-6294-4a13-9ac1-9783190fa794', '', 'Invoice', 'Reject', 'e4a12657-a23b-48b1-98a8-06f7bdaa0f7f', 'FeedbackSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('dd12a7d2-b143-429b-9bc7-f90273a4a6ef', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '00000000-0000-0000-0000-000000000000', '{"RootProof":"0e231249f13400e4bedacc41eb9930a8","Nodes":[{"SyncTreeNodeID":"","Value":"InvoiceID:Inv123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"Currency:USD","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"Amount:6000","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"7f1b22cdb9ccda7de6a3bb94d21e3189","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"b70da7abea1d8f05699ef195d2b04c6e","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"0e231249f13400e4bedacc41eb9930a8","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'NewVersion', '0e231249f13400e4bedacc41eb9930a8', '58f0dc4a-99bf-4aa1-8205-226c48a8e729', '58f0dc4a-99bf-4aa1-8205-226c48a8e729', '94dd403e-e1e6-4834-826a-5a87a471b44e', 'ef9ea07d-6294-4a13-9ac1-9783190fa794', 'NewVersion suggested', 'Invoice', 'Suggest', '5c4f4f57-ac12-4ae8-a8d3-f290870f1aee', 'SuggestionSent');
INSERT INTO public.offchain_process_messages (id, sender_id, receiver_id, topic, referenced_offchain_process_message_id, baseledger_sync_tree_json, workstep_type, business_object_proof, tendermint_transaction_id_of_stored_proof, baseledger_transaction_id_of_stored_proof, baseledger_business_object_id, referenced_baseledger_business_object_id, status_text_message, business_object_type, baseledger_transaction_type, referenced_baseledger_transaction_id, entry_type) VALUES ('393d82ea-a99e-494a-adcb-65f8fdd1f29f', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'dd12a7d2-b143-429b-9bc7-f90273a4a6ef', '{"RootProof":"0e231249f13400e4bedacc41eb9930a8","Nodes":[{"SyncTreeNodeID":"","Value":"InvoiceID:Inv123","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":0},{"SyncTreeNodeID":"","Value":"Currency:USD","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":1},{"SyncTreeNodeID":"","Value":"Amount:6000","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":2},{"SyncTreeNodeID":"","Value":"","IsLeaf":true,"IsRoot":false,"IsHash":false,"IsCovered":false,"Level":0,"Index":3},{"SyncTreeNodeID":"","Value":"7f1b22cdb9ccda7de6a3bb94d21e3189","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":0},{"SyncTreeNodeID":"","Value":"b70da7abea1d8f05699ef195d2b04c6e","IsLeaf":false,"IsRoot":false,"IsHash":true,"IsCovered":false,"Level":1,"Index":1},{"SyncTreeNodeID":"","Value":"0e231249f13400e4bedacc41eb9930a8","IsLeaf":false,"IsRoot":true,"IsHash":true,"IsCovered":false,"Level":2,"Index":0}]}', 'Feedback', '', '01f9a344-9937-410e-b415-455105c4567b', '01f9a344-9937-410e-b415-455105c4567b', '00000000-0000-0000-0000-000000000000', '94dd403e-e1e6-4834-826a-5a87a471b44e', '', 'Invoice', 'Approve', '58f0dc4a-99bf-4aa1-8205-226c48a8e729', 'FeedbackSent');

INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('f7c1e3e7-fa89-4e32-97df-21a8f4c1b508', '10375', '467dca32-ba2a-4352-9685-6144df7336e4', '2021-07-21 23:57:19.333613+00', 'SuggestionSent', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Initial', 'Suggest', '467dca32-ba2a-4352-9685-6144df7336e4', '00000000-0000-0000-0000-000000000000', 'PurchaseOrder', 'b1ade3ef-76aa-45c0-9c08-8c9c295f3601', '00000000-0000-0000-0000-000000000000', '6a55d674-3ac9-47a3-a83e-b354c623936f', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '48C91DCD2C506AD91A9A5A9F892CB987BAD4F2AB1B4BD9C755081415B937248D', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('579b2c3e-08b9-433e-bed2-25fc63319c41', '10443', 'c0f5591b-fbf7-483b-bf1e-c18e362c7a06', '2021-07-21 23:58:29.885871+00', 'FeedbackSent', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Reject', 'c0f5591b-fbf7-483b-bf1e-c18e362c7a06', '467dca32-ba2a-4352-9685-6144df7336e4', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', 'b1ade3ef-76aa-45c0-9c08-8c9c295f3601', '89f90a49-78c2-4dc9-952a-1c130d5e75f5', '6a55d674-3ac9-47a3-a83e-b354c623936f', 'COMMITTED', '786F31A41EB3265BD5E396C381F20FE7203295743881E34E837011A8227875C2', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('e24db1e8-f753-473d-8641-05a949e4f759', '10495', '7fc8d994-0c3a-4cea-9946-828d5d60a99e', '2021-07-21 23:59:25.051741+00', 'SuggestionSent', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'NewVersion', 'Suggest', '7fc8d994-0c3a-4cea-9946-828d5d60a99e', 'c0f5591b-fbf7-483b-bf1e-c18e362c7a06', 'PurchaseOrder', 'beda5b1f-0e6d-4cd6-b960-f9c0229e1d47', 'b1ade3ef-76aa-45c0-9c08-8c9c295f3601', 'e3584552-7582-4b21-814a-bd8fe0fe7e39', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '4E305BDD61C8F9B7A7493EED492F2961577B62894E0681634E47F97FFCC663FE', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('5d0caf35-25a5-4f2c-bc74-4a63e6db77b8', '10556', 'ce7f46b0-a1bc-48be-8b4a-95affab4a89b', '2021-07-22 00:00:24.178021+00', 'FeedbackSent', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Approve', 'ce7f46b0-a1bc-48be-8b4a-95affab4a89b', '7fc8d994-0c3a-4cea-9946-828d5d60a99e', 'PurchaseOrder', '00000000-0000-0000-0000-000000000000', 'beda5b1f-0e6d-4cd6-b960-f9c0229e1d47', '6a03eca3-e92c-48c8-a55b-2f3a3794495c', 'e3584552-7582-4b21-814a-bd8fe0fe7e39', 'COMMITTED', '6BF6E13D591A5E461E3FBDAB029F3B2FEA905E7A0B7071E78378B5983C37835F', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('d0cbbeb3-cedc-4a18-a188-1abbbf365f51', '10689', 'e4a12657-a23b-48b1-98a8-06f7bdaa0f7f', '2021-07-22 00:02:39.820183+00', 'SuggestionSent', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'NextWorkstep', 'Suggest', 'e4a12657-a23b-48b1-98a8-06f7bdaa0f7f', 'ce7f46b0-a1bc-48be-8b4a-95affab4a89b', 'Invoice', 'ef9ea07d-6294-4a13-9ac1-9783190fa794', 'beda5b1f-0e6d-4cd6-b960-f9c0229e1d47', '05aa6e2b-5af4-49ba-9fac-07ab442023a8', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '119A898EB38AFB7A377816B56834E98B26AA2915C8A1E4BB7A2C07A3795E6EAC', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('0a81bc74-8c69-427a-bb07-e7d136d09688', '10753', '5c4f4f57-ac12-4ae8-a8d3-f290870f1aee', '2021-07-22 00:03:45.05249+00', 'FeedbackSent', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Reject', '5c4f4f57-ac12-4ae8-a8d3-f290870f1aee', 'e4a12657-a23b-48b1-98a8-06f7bdaa0f7f', 'Invoice', '00000000-0000-0000-0000-000000000000', 'ef9ea07d-6294-4a13-9ac1-9783190fa794', '52c1bfb1-bc1b-417a-a204-160cc02a56ec', '05aa6e2b-5af4-49ba-9fac-07ab442023a8', 'COMMITTED', '5F76A1D22AB3CAD10E9ABDCFC47B9301776B92F2A62041D2F663F76B4DD42895', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('863ac1ac-d1c8-4d3c-9c9a-ac81dbfb8284', '10865', '58f0dc4a-99bf-4aa1-8205-226c48a8e729', '2021-07-22 00:05:40.257104+00', 'SuggestionSent', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'NewVersion', 'Suggest', '58f0dc4a-99bf-4aa1-8205-226c48a8e729', '5c4f4f57-ac12-4ae8-a8d3-f290870f1aee', 'Invoice', '94dd403e-e1e6-4834-826a-5a87a471b44e', 'ef9ea07d-6294-4a13-9ac1-9783190fa794', 'dd12a7d2-b143-429b-9bc7-f90273a4a6ef', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '76D455B9F743E940BEADBFCE807FC01891BA907072C7AEF5664ADBDE6533D0BB', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('3f704070-5d57-44fa-b5ae-147505c73d6a', '10898', '01f9a344-9937-410e-b415-455105c4567b', '2021-07-22 00:06:14.872003+00', 'FeedbackSent', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Approve', '01f9a344-9937-410e-b415-455105c4567b', '58f0dc4a-99bf-4aa1-8205-226c48a8e729', 'Invoice', '00000000-0000-0000-0000-000000000000', '94dd403e-e1e6-4834-826a-5a87a471b44e', '393d82ea-a99e-494a-adcb-65f8fdd1f29f', 'dd12a7d2-b143-429b-9bc7-f90273a4a6ef', 'COMMITTED', '78EF517142889FF772A5E23222768DAC2E537068FC6E9187EDB16F34F45F297A', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('b15607d6-8375-4b3e-8cf4-ea5c7cc3c74c', '11023', 'd130279e-6373-4744-82f9-fc542183eadc', '2021-07-22 00:08:20.176064+00', 'SuggestionSent', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'FinalWorkstep', 'Suggest', 'd130279e-6373-4744-82f9-fc542183eadc', '01f9a344-9937-410e-b415-455105c4567b', 'MerkleTreeForExiting', '12794521-4d92-40f6-b28f-c7baa282a7c7', '94dd403e-e1e6-4834-826a-5a87a471b44e', 'd61a4a37-6ab6-40da-a9c2-5af3024014ca', '00000000-0000-0000-0000-000000000000', 'COMMITTED', '1E225F8837204D4EA48A34052BACC47197D291052CA2A0BBB2F794E468A58171', '61967691-5be4-425f-b45f-d83ccff27e7f');
INSERT INTO public.trustmesh_entries (id, tendermint_block_id, tendermint_transaction_id, tendermint_transaction_timestamp, entry_type, sender_org_id, receiver_org_id, workgroup_id, workstep_type, baseledger_transaction_type, baseledger_transaction_id, referenced_baseledger_transaction_id, business_object_type, baseledger_business_object_id, referenced_baseledger_business_object_id, offchain_process_message_id, referenced_process_message_id, commitment_state, transaction_hash, trustmesh_id) VALUES ('d1a3b876-c877-4183-94b3-e8528585f4e0', '11067', '83b858a8-8cbc-4f54-ac4b-422b5c9ccc5e', '2021-07-22 00:09:05.007649+00', 'FeedbackSent', '61ded832-b7ca-4100-8bc1-fb0935ff4436', '68f6fb46-7fe5-4536-ac98-52b475418f7e', '68f6fb46-7fe5-4536-ac98-52b475418f7e', 'Feedback', 'Approve', '83b858a8-8cbc-4f54-ac4b-422b5c9ccc5e', 'd130279e-6373-4744-82f9-fc542183eadc', 'MerkleTreeForExiting', '00000000-0000-0000-0000-000000000000', '12794521-4d92-40f6-b28f-c7baa282a7c7', '118118dd-ff64-439f-b0bf-f02b443654d2', 'd61a4a37-6ab6-40da-a9c2-5af3024014ca', 'COMMITTED', '90C0A049FDEFCE154E476FD1E012B1BB6506E31136A253C435127675C184E7A4', '61967691-5be4-425f-b45f-d83ccff27e7f');