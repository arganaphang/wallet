CREATE TYPE TransactionType AS ENUM ('income', 'outcome');

CREATE TABLE IF NOT EXISTS "public"."transactions" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "amount" BIGINT NOT NULL,
  "category" VARCHAR NOT NULL,
  "type" TransactionType NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
