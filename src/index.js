const app = require('./app');
require('dotenv').config();
const logger = require('./logger');
const port = process.env.PORT || 5000;
app.listen(port, () => {
    /* eslint-disable no-console */
    logger.info(`Listening: http://localhost:${port}`);
    /* eslint-enable no-console */
});