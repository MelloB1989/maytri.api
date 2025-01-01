CREATE TABLE IF NOT EXISTS "messages" (
	"id" varchar PRIMARY KEY NOT NULL,
	"chat_id" varchar,
	"sender_id" varchar,
	"message" text,
	"created_at" timestamp,
	CONSTRAINT "messages_id_unique" UNIQUE("id")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "messages" ADD CONSTRAINT "messages_chat_id_chats_id_fk" FOREIGN KEY ("chat_id") REFERENCES "public"."chats"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
