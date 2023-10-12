ALTER TABLE "tickets" DROP FOREIGN KEY IF EXISTS fk_tickets_event_id;
ALTER TABLE "events" DROP FOREIGN KEY IF EXISTS fk_events_user_id;
ALTER TABLE "events" DROP FOREIGN KEY IF EXISTS fk_events_image_id;

DROP TABLE IF EXISTS "tickets";
DROP TABLE IF EXISTS "images";
DROP TABLE IF EXISTS "events";
DROP TABLE IF EXISTS "users";

DROP TYPE IF EXISTS "stock_type";
DROP TYPE IF EXISTS "access_type";
DROP TYPE IF EXISTS "ticket_type";
DROP TYPE IF EXISTS "event_type";
DROP TYPE IF EXISTS "recurrence_type";

