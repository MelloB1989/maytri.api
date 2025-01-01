import {
  pgTable,
  json,
  varchar,
  timestamp,
  unique,
  text,
  integer,
  boolean,
} from "drizzle-orm/pg-core";
import { sql } from "drizzle-orm";

export const users = pgTable("users", {
  id: varchar("id").primaryKey().unique(),
  email: varchar("email").default(""),
  phone: varchar("phone").unique(),
  address: varchar("address"),
  password: varchar("password"),
  name: varchar("name"),
  profile_image: varchar("profile_image"),
  age: integer("age"),
  location: varchar("location"),
});

export const maytri = pgTable("maytri", {
  id: varchar("id").primaryKey().unique(),
  user_id: varchar("user_id").references(() => users.id),
  description: varchar("description"),
  image: varchar("image"),
  age: integer("age"),
  gender: varchar("gender"),
  profession: varchar("profession"),
  created_at: timestamp("created_at"),
});

export const chats = pgTable("chats", {
  id: varchar("id").primaryKey().unique(),
  user_id: varchar("user_id").references(() => users.id),
  type: varchar("type"),
  created_at: timestamp("created_at"),
});

export const chatParticipants = pgTable("chat_participants", {
  id: varchar("id").primaryKey().unique(),
  chat_id: varchar("chat_id").references(() => chats.id),
  participant_id: varchar("participant_id"),
});

export const groupDetails = pgTable("group_details", {
  chat_id: varchar("chat_id").references(() => chats.id),
  group_name: varchar("group_name"),
  group_image: varchar("group_image"),
  group_description: varchar("group_description"),
});

export const messages = pgTable("messages", {
  id: varchar("id").primaryKey().unique(),
  chat_id: varchar("chat_id").references(() => chats.id),
  sender_id: varchar("sender_id"),
  message: text("message"),
  created_at: timestamp("created_at"),
});
