ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_pkey;

ALTER TABLE ONLY public.trustmesh_entries DROP CONSTRAINT trustmesh_entries_offchain_process_message_id_offchain_process_messages_id_foreign;

ALTER TABLE ONLY public.trustmesh_entries
    ALTER COLUMN offchain_process_message_id TYPE text;
