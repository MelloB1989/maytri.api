import type { Config } from "drizzle-kit";

export default {
  schema: "./db/schema.ts",
  out: "./drizzle",
  dialect: "postgresql", // 'pg' | 'mysql2' | 'better-sqlite' | 'libsql' | 'turso'
  dbCredentials: {
    host: "100.99.161.86",
    user: "mellob",
    password: "mellob1989",
    database: "maytri",
    ssl: false,
  },
} satisfies Config;
