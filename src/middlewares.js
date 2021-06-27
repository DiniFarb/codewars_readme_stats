const logger = require('./logger');

function notFound(req, res, next) {
    res.status(404);
    const error = new Error(`🔍 - Not Found - ${req.originalUrl}`);
    next(error);
}

/* eslint-disable no-unused-vars */
function errorHandler(err, req, res, next) {
    /* eslint-enable no-unused-vars */
    logger.error(err.message)
    const statusCode = res.statusCode !== 200 ? res.statusCode : 500;
    res.status(statusCode);
    res.send(err.message);
}

module.exports = {
    notFound,
    errorHandler
};