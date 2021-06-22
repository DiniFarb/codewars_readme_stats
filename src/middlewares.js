const logger = require('./logger');

function notFound(req, res, next) {
  res.status(404);
  const error = new Error(`🔍 - Not Found - ${req.originalUrl}`);
  logger.error(error)
  next(error);
}

/* eslint-disable no-unused-vars */
function errorHandler(err, req, res, next) {
  /* eslint-enable no-unused-vars */
  logger.error(err)
  const statusCode = res.statusCode !== 200 ? res.statusCode : 500;
  res.status(statusCode);
  res.json({
    message: err.message,
    stack: process.env.NODE_ENV === 'production' ? '🥞' : err.stack
  });
}

module.exports = {
  notFound,
  errorHandler
};
