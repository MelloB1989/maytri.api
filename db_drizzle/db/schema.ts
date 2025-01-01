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
