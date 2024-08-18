BEGIN;

CREATE TABLE "account" (
    "createdAt" timestamptz,
    "updatedAt" timestamptz,
    "deletedAt" timestamptz,
    "id" UUID,
    "userId" UUID NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    PRIMARY KEY ("id")
);

COMMIT;