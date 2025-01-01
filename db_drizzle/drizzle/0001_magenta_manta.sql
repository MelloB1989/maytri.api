CREATE TABLE IF NOT EXISTS "chat_participants" (
	"id" varchar PRIMARY KEY NOT NULL,
	"chat_id" varchar,
	"user_id" varchar,
	"maytri_id" varchar,
	CONSTRAINT "chat_participants_id_unique" UNIQUE("id")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "chats" (
	"id" varchar PRIMARY KEY NOT NULL,
	"user_id" varchar,
	"type" varchar,
	"created_at" timestamp,
	CONSTRAINT "chats_id_unique" UNIQUE("id")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "group_details" (
	"chat_id" varchar,
	"group_name" varchar,
	"group_image" varchar,
	"group_description" varchar
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "chat_participants" ADD CONSTRAINT "chat_participants_chat_id_chats_id_fk" FOREIGN KEY ("chat_id") REFERENCES "public"."chats"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "chat_participants" ADD CONSTRAINT "chat_participants_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "chat_participants" ADD CONSTRAINT "chat_participants_maytri_id_maytri_id_fk" FOREIGN KEY ("maytri_id") REFERENCES "public"."maytri"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "chats" ADD CONSTRAINT "chats_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "group_details" ADD CONSTRAINT "group_details_chat_id_chats_id_fk" FOREIGN KEY ("chat_id") REFERENCES "public"."chats"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
