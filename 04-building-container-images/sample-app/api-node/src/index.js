const { getDateTime } = require('./db');

const express = require('express');

const app = express();
const port = 3000;

app.get('/', async (req, res) => {
  const dateTime = await getDateTime();
  res.send(dateTime);
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
