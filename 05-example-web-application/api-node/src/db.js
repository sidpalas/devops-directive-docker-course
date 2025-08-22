const { Pool } = require("pg");
const { DateTime } = require("luxon");
const fs = require("fs");

const databaseUrl =
  process.env.DATABASE_URL ||
  fs.readFileSync(process.env.DATABASE_URL_FILE, "utf8");

const pool = new Pool({ connectionString: databaseUrl });

pool.on("error", (err) => {
  console.error("Unexpected error on idle client", err);
  process.exit(-1);
});

const getDateTime = async (clientTimeZone = "Africa/Harare") => {
  const client = await pool.connect();
  try {
    const res = await client.query("SELECT NOW() as now;");
    const utcTime = res.rows[0].now;
    const localTime = DateTime.fromJSDate(utcTime).setZone(clientTimeZone);
    return {
      api: "node",
      now: localTime.toISO({
        includeOffset: true,
        suppressMilliseconds: false,
      }),
    };
  } catch (err) {
    console.error(err.stack);
    return { error: "Failed to fetch time" };
  } finally {
    client.release();
  }
};

module.exports = { getDateTime };
