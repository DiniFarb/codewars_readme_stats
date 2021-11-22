import pkg from 'winston';
const { createLogger, transports, format } = pkg;

const logger = createLogger({
    transports: [
        new transports.Console({
            format: format.printf(info => `${info.level}: ${info.message}`)
        }),
    ]
});
export default logger