import { SeedPg } from "@snaplet/seed/adapter-pg";
import { defineConfig } from "@snaplet/seed/config";
import { Client } from "pg";

export default defineConfig({
  adapter: async () => {
    const client = new Client(
      'postgres://postgres.ocxtbfagmkbrozrmaygm:fqaH3Xhm6xwdKmvV@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres',
    );
    await client.connect();
    return new SeedPg(client);
  },
});