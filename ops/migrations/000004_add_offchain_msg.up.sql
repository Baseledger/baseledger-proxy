CREATE TABLE public.offchain_process_messages (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    sender_id                         text,
	receiver_id                       text,
	topic                            text,
	referenced_offchain_process_message_id text,
	business_object                       text,
	workstep_type                         text,
	hash                                 text,
	tendermint_transaction_id_of_stored_proof text,
	baseledger_transaction_id_of_stored_proof text,
    baseledger_business_object_id           text,
	referenced_baseledger_business_object_id text,
	status_text_message                    text
);

ALTER TABLE public.offchain_process_messages OWNER TO baseledger;

ALTER TABLE ONLY public.offchain_process_messages ADD CONSTRAINT offchain_process_messagess_pkey PRIMARY KEY (id);