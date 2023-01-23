const { getDateTime } = require('./db');

const express = require('express');

const app = express();
const port = 3000;

app.get('/', async (req, res) => {
  const dateTime = await getDateTime();
  const response = dateTime;
  response.api = 'node';
  res.send(response);
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
