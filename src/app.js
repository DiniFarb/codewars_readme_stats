const express = require('express');
const morgan = require('morgan');
const helmet = require('helmet');
const axios = require('axios');
const cors = require('cors');
require('dotenv').config();
const middlewares = require('./middlewares');
const app = express();
const logger = require('./logger');

app.use(helmet());
app.use(morgan('tiny'));
app.use(cors());
app.use(express.json());

/**
 * Rely to codewars user infos
 */
app.get('/codewars', async(req, res, next) => {
    const codewars_Url = 'https://www.codewars.com/api/v1/users/';
    try {
    const { data } = await axios.get(`${codewars_Url}/${req.query.user}`);
    const answer = `
         <svg
        width="495"
        height="195"
        viewBox="0 0 495 195"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >${data}</svg>`;
    res.send(answer);    
    } catch(err) {
       next(err); 
    }
});

app.use(middlewares.notFound);
app.use(middlewares.errorHandler);

module.exports = app;