DROP TABLE public.trustmesh_entries;
DROP TABLE public.workgroup_members;
DROP TABLE public.workgroups;
DROP TABLE public.organizations;

ALTER TABLE ONLY public.offchain_process_messages DROP CONSTRAINT offchain_process_messagess_pkey;

DROP TABLE public.offchain_process_messages;

ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_pkey;

ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_offchain_process_message_id_offchain_process_messages_id_foreign;

ALTER TABLE ONLY public.trustmesh_entries
    ALTER COLUMN offchain_process_message_id TYPE text;
