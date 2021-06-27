const express = require('express');
const helmet = require('helmet');
const cors = require('cors');
const rateLimit = require("express-rate-limit");
require('dotenv').config();
const middlewares = require('./middlewares');
const logger = require('./logger');
const codewars = require('./routes/codewars');

const app = express();
app.use(helmet());
app.use(cors());
app.set('trust proxy', 1);

const limiter = rateLimit({
    windowMs: 60 * 1000,
    max: 20
});

app.use(limiter);

app.get('/', async(req, res) => {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    logger.info(`redirect: ${ip}`);
    res.redirect('https://github.com/andreasvogt89/codewars_api');
});

app.use('/codewars', codewars);

app.use(middlewares.notFound);
app.use(middlewares.errorHandler);

module.exports = app;