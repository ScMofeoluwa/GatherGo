CREATE TYPE "recurrence_type" AS ENUM (
  'single',
  'daily',
  'weekly',
  'biweekly'
);

CREATE TYPE "event_type" AS ENUM (
  'live',
  'online'
);

CREATE TYPE "ticket_type" AS ENUM (
  'single',
  'group'
);

CREATE TYPE "access_type" AS ENUM (
  'free',
  'paid'
);

CREATE TYPE "stock_type" AS ENUM (
  'limited',
  'unlimited'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar UNIQUE,
  "password" varchar NOT NULL,
  "verified" boolean DEFAULT false,
  "registered_at" timestamptz DEFAULT now()
);

CREATE TABLE "events" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "type" event_type NOT NULL,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "url_slug" varchar UNIQUE,
  "category" varchar NOT NULL,
  "recurrence" recurrence_type NOT NULL,
  "timezone" varchar NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "start_time" time NOT NULL,
  "end_time" time NOT NULL,
  "social_details" jsonb,
  "location" geography,
  "image_id" uuid
);

CREATE TABLE "tickets" (
  "id" uuid PRIMARY KEY,
  "event_id" uuid NOT NULL,
  "invite" boolean DEFAULT false,
  "type" ticket_type NOT NULL,
  "access_type" access_type NOT NULL,
  "name" varchar NOT NULL,
  "stock_type" stock_type DEFAULT 'unlimited',
  "stock_count" int,
  "purchase_limit" int,
  "group_size" int,
  "description" varchar NOT NULL,
  "price" int
);

CREATE TABLE "images" (
  "id" uuid PRIMARY KEY,
  "event_id" uuid NOT NULL,
  "url" varchar NOT NULL
);

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "images" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE;
