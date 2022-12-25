-- Table Definition
CREATE TABLE "public"."record"
(
    "id"         int4 NOT NULL DEFAULT nextval('record_id_seq'::regclass),
    "book"       text NOT NULL,
    "created_at" timestamp,
    "updated_at" timestamp,
    PRIMARY KEY ("id")
);