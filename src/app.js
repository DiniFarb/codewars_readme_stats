import express from 'express';
import helmet from 'helmet';
import cors from 'cors';
import rateLimit from 'express-rate-limit';
import { notFound, errorHandler } from './middlewares.js';
import logger from './logger.js';
import codewars from './routes/codewars.js';

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

app.use(notFound);
app.use(errorHandler);

export default app;