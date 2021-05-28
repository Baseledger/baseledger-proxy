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

CREATE TABLE public.trustmeshes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    tendermintBlockId text,
    tendermintTransactionId text,
    tendermintTransactionTimestamp timestamp with time zone,

    sender text,
    receiver text,
    workgroupId text,

    workstepType text,
    baseledgerTransactionType text,

    baseledgerTransactionId text,
    referencedBaseledgerTransactionId text,

    businessObjectType text,
    baseledgerBusinessObjectID text,
    referencedBaseledgerBusinessObjectID text,

    offchainProcessMessageID text,
    referencedProcessMessageID text
);

ALTER TABLE public.trustmeshes OWNER TO baseledger;