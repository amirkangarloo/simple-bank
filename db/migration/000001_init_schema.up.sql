CREATE TABLE "account" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "owner" varchar NOT NULL,
    "balance" bigint NOT NULL,
    "currency" varchar NOT NULL
);

CREATE TABLE "entries" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "account_id" bigint NOT NULL,
    "amount" bigint NOT NULL
);

CREATE TABLE "transfers" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "from_account_id" bigint NOT NULL,
    "to_account_id" bigint NOT NULL,
    "amount" bigint NOT NULL
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ( "from_account_id", "to_account_id" );

COMMENT ON COLUMN "entries"."amount" IS 'It can be positive or negative';

COMMENT ON COLUMN "transfers"."amount" IS 'It can be positive';

ALTER TABLE "entries"
ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfers"
ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfers"
ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");