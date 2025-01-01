CREATE TABLE IF NOT EXISTS "maytri" (
	"id" varchar PRIMARY KEY NOT NULL,
	"user_id" varchar,
	"description" varchar,
	"image" varchar,
	"age" integer,
	"gender" varchar,
	"profession" varchar,
	"created_at" timestamp,
	CONSTRAINT "maytri_id_unique" UNIQUE("id")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "users" (
	"id" varchar PRIMARY KEY NOT NULL,
	"email" varchar DEFAULT '',
	"phone" varchar,
	"address" varchar,
	"password" varchar,
	"name" varchar,
	"profile_image" varchar,
	"age" integer,
	"location" varchar,
	CONSTRAINT "users_id_unique" UNIQUE("id"),
	CONSTRAINT "users_phone_unique" UNIQUE("phone")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "maytri" ADD CONSTRAINT "maytri_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
