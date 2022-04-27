-- Table: public.pets

-- DROP TABLE IF EXISTS public.pets;

CREATE TABLE IF NOT EXISTS public.pets
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    gender character(1) COLLATE pg_catalog."default",
    date_of_birth date,
    status boolean NOT NULL DEFAULT true,
    deleted_at date,
    created_at date,
    updated_at date,
    specie character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT pet_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pets
    OWNER to postgres;