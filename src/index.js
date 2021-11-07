import app from "./app.js";
import logger from "./logger.js";

const port = process.env.PORT || 5000;
app.listen(port, () => {
  logger.info(`service is running on:: [${port}]`);
});