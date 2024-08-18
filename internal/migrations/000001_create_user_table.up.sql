BEGIN;

CREATE TABLE "user" (
    "createdAt" timestamptz,
    "updatedAt" timestamptz,
    "deletedAt" timestamptz,
    "id" UUID,
    "email" VARCHAR(50) NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "providers" text[] NOT NULL,
    "refreshToken" VARCHAR(255),
    PRIMARY KEY ("id"),
    CONSTRAINT "uniq_email" UNIQUE ("email")
);

COMMIT;