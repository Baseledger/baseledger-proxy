ALTER TABLE ONLY public.trustmesh_entries ADD CONSTRAINT trustmesh_entries_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trustmesh_entries
    ALTER COLUMN offchain_process_message_id TYPE uuid USING (uuid_generate_v4());

ALTER TABLE ONLY public.trustmesh_entries
    ADD CONSTRAINT trustmesh_entries_offchain_process_message_id_offchain_process_messages_id_foreign FOREIGN KEY (offchain_process_message_id) REFERENCES public.offchain_process_messages(id) ON UPDATE CASCADE ON DELETE CASCADE;