CREATE SEQUENCE "company_id_seq";
SELECT setval('company_id_seq', (SELECT MAX("id") FROM "company"));
ALTER TABLE "company"
    ALTER COLUMN "id" SET DEFAULT nextval('company_id_seq'),
    ALTER COLUMN "id" SET NOT NULL;