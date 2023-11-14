-- CreateTable
CREATE TABLE "transactions" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "amount" REAL NOT NULL,
    "date" TEXT NOT NULL,
    "account_id" TEXT NOT NULL
);
