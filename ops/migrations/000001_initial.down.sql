ALTER TABLE ONLY public.offchain_process_messages DROP CONSTRAINT offchain_process_messagess_pkey;
ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_pkey;
ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_offchain_process_message_id_offchain_process_messages_id_foreign;

ALTER TABLE ONLY public.trustmesh_entries ALTER COLUMN offchain_process_message_id TYPE text;
ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_trustmesh_id_trustmeshes_id_foreign;
ALTER TABLE ONLY public.trustmeshes DROP CONSTRAINT trustmeshes_pkey;

DROP FUNCTION IF EXISTS set_trustmesh_entry_group CASCADE;
DROP TRIGGER IF EXISTS trustmesh_entry_insert_trigger ON trustmesh_entries;

DROP TABLE public.trustmesh_entries;
DROP TABLE public.offchain_process_messages;
DROP TABLE public.workgroup_members;
DROP TABLE public.workgroups;
DROP TABLE public.organizations;
DROP TABLE public.trustmeshes;