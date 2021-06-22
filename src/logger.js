const { createLogger, transports, format } = require('winston');

const logger = createLogger({
    format: format.combine(
        format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss:ms' }),
        format.printf(info => `${info.timestamp} ${info.level}: ${info.message}`)
    ),
    transports: [
        new transports.File({
            filename: './logs/ServerLog.log',
            json: false,
            maxsize: 5242880,
            maxFiles: 1,
        }),
        new transports.Console(),
    ]
});

module.exports = logger;