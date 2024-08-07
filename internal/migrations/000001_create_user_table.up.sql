BEGIN;

CREATE TABLE "user" (
    "createdAt" timestamptz,
    "updatedAt" timestamptz,
    "deletedAt" timestamptz,
    "id" VARCHAR(16),
    "email" VARCHAR(50),
    "password" VARCHAR(255),
    "name" VARCHAR(50),
    "providers" jsonb,
    "refreshToken" VARCHAR(255),
    PRIMARY KEY ("id"),
    CONSTRAINT "uniq_email" UNIQUE ("email")
);

COMMIT;