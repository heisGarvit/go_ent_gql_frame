-- Create "companies" table
CREATE TABLE "companies" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "company_id" uuid NULL, "name" character varying NOT NULL, "is_disabled" boolean NOT NULL DEFAULT false, "owner_id" uuid NULL, PRIMARY KEY ("id"));
-- Create index "company_company_id" to table: "companies"
CREATE INDEX "company_company_id" ON "companies" ("company_id");
-- Create index "company_created_at" to table: "companies"
CREATE INDEX "company_created_at" ON "companies" ("created_at");
-- Create "permissions" table
CREATE TABLE "permissions" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "company_id" uuid NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "read" boolean NOT NULL DEFAULT false, "write" boolean NOT NULL DEFAULT false, "patch" boolean NOT NULL DEFAULT false, "delete" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- Create index "permission_company_id" to table: "permissions"
CREATE INDEX "permission_company_id" ON "permissions" ("company_id");
-- Create index "permission_created_at" to table: "permissions"
CREATE INDEX "permission_created_at" ON "permissions" ("created_at");
-- Create "roles" table
CREATE TABLE "roles" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "name" character varying NOT NULL, "description" character varying NOT NULL, "company_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "roles_companies_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "role_company_id" to table: "roles"
CREATE INDEX "role_company_id" ON "roles" ("company_id");
-- Create index "role_created_at" to table: "roles"
CREATE INDEX "role_created_at" ON "roles" ("created_at");
-- Create "role_permissions" table
CREATE TABLE "role_permissions" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "company_id" uuid NULL, "role_id" uuid NULL, "permission_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "role_permissions_companies_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "role_permissions_permissions_permission" FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "role_permissions_roles_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "rolepermission_company_id" to table: "role_permissions"
CREATE INDEX "rolepermission_company_id" ON "role_permissions" ("company_id");
-- Create index "rolepermission_created_at" to table: "role_permissions"
CREATE INDEX "rolepermission_created_at" ON "role_permissions" ("created_at");
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "name" character varying NULL, "email" character varying NOT NULL, "mobile" character varying NOT NULL, "password" character varying NOT NULL, "is_disabled" boolean NOT NULL DEFAULT false, "role_id" uuid NULL, "company_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "users_companies_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "user_company_id" to table: "users"
CREATE INDEX "user_company_id" ON "users" ("company_id");
-- Create index "user_company_id_email" to table: "users"
CREATE UNIQUE INDEX "user_company_id_email" ON "users" ("company_id", "email");
-- Create index "user_company_id_mobile" to table: "users"
CREATE UNIQUE INDEX "user_company_id_mobile" ON "users" ("company_id", "mobile");
-- Create index "user_created_at" to table: "users"
CREATE INDEX "user_created_at" ON "users" ("created_at");
