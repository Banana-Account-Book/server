BEGIN;

CREATE TABLE "transaction" (
    "createdAt" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" timestamptz,
    "id" UUID PRIMARY KEY,
    "userId" UUID NOT NULL,
    "accountBookId" UUID NOT NULL,
    "title" VARCHAR(50) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "registeredAt" timestamptz NOT NULL,
    "periodStartOn" DATE NOT NULL,
    "periodEndOn" DATE NULL,
    "type" VARCHAR(20) NOT NULL,
    "amount" INT NOT NULL,
    "repeatType" VARCHAR(20) NOT NULL
);

CREATE TABLE "exclusive" (
    "id" SERIAL PRIMARY KEY,
    "userId" UUID NOT NULL,
    "periodStartOn" DATE NOT NULL,
    "periodEndOn" DATE,
    "title" VARCHAR(50) NOT NULL,
    "description" VARCHAR(255),
    "amount" INTEGER NOT NULL,
    "transactionId" UUID NOT NULL,
    FOREIGN KEY ("transactionId") REFERENCES "transaction"(id)
);

COMMIT;