BEGIN;

CREATE TABLE "role" (
    "createdAt" timestamptz,
    "updatedAt" timestamptz,
    "deletedAt" timestamptz,
    "id" SERIAL PRIMARY KEY,
    "userId" UUID NOT NULL,
    "accountBookId" UUID NOT NULL,
    "type" VARCHAR(20) NOT NULL
);

COMMIT;