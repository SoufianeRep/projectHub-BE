CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "last_signin" timestamp DEFAULT (now()),
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "teams" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "projects" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "language" varchar,
  "length" int,
  "transcript" varchar,
  "team_id" bigint UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "role" varchar NOT NULL,
  "user_id" bigint UNIQUE NOT NULL,
  "team_id" bigint UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "projects" ("team_id");

CREATE INDEX ON "roles" ("user_id");

CREATE INDEX ON "roles" ("team_id");

ALTER TABLE "projects" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");

ALTER TABLE "roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "roles" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");
