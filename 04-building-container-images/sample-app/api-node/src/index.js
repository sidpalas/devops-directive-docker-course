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

const server = app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});

process.on('SIGTERM', () => {
  debug('SIGTERM signal received: closing HTTP server');
  server.close(() => {
    debug('HTTP server closed');
  });
});
