"use strict";

const PORT = 9000;
const express = require('express');
const app = express();

app.use(express.static('dist'));

app.listen(PORT, function () {
  console.log(`Example app listening on port ${PORT}!`);
});
