-- Table Definition
CREATE TABLE "public"."book"
(
    "book"  text NOT NULL,
    "price" int8 DEFAULT 0,
    PRIMARY KEY ("book")
);