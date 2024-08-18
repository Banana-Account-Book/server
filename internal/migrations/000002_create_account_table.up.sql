BEGIN;

CREATE TABLE "account" (
    "createdAt" timestamptz,
    "updatedAt" timestamptz,
    "deletedAt" timestamptz,
    "id" UUID,
    "userId" UUID,
    "name" VARCHAR(50),
    PRIMARY KEY ("id")
);

COMMIT;