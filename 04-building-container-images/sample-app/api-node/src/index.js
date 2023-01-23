const { getDateTime } = require('./db');

const express = require('express');
const cors = require('cors');

const app = express();
app.use(
  cors({
    origin: 'http://127.0.0.1:5173',
  })
);
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
