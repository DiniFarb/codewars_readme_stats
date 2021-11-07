import express from "express";
import logger from "../logger.js";
import axios from "axios";
import fs from "fs";
import path from "path";
import sIcons from "simple-icons";
import specialIcons from "../utils/icon-mapper.js";
const router = express();
const __dirname = path.dirname(new URL(import.meta.url).pathname);

const LevelColors = {
  1: "#866CC7",
  2: "#866CC7",
  3: "#3C7EBB",
  4: "#3C7EBB",
  5: "#ECB613",
  6: "#ECB613",
  7: "#E6E6E6",
  8: "#E6E6E6",
  dan: "#999999",
};

let top_languages_template = (icons) => `
<g transform="translate(150, 190)">
<g class="stats" style="animation-delay: 1050ms">      
  <text class="stat bold"  y="12.5">Top trained languages</text>
  <text 
    class="stat" 
    x="170" 
    y="12.5" 
  >
  </text>
  ${icons}
  </g>
</g>`;

let icon_template = (x) => `
<g transform="translate(${x},20)">
  {svg}
</g>`;

let no_icon_found_template = (iconName) => `
<svg viewBox="0 0 150 150" class="fail-icon-text">
  <text x="10" y="10" alignment-baseline="central" dominant-baseline="central" text-anchor="middle">
  ${iconName}
  </text>
</svg>`;

router.get("/", async (req, res, next) => {
  const ip = req.headers["x-forwarded-for"] || req.socket.remoteAddress;
  logger.info(`crad request: ${ip}`);
  const codewars_Url = "https://www.codewars.com/api/v1/users";
  try {
    const { data } = await axios.get(`${codewars_Url}/${req.query.user}`);
    res.setHeader("Content-Type", "image/svg+xml");
    res.setHeader("Cache-Control", `public, max-age=7200`);
    const levelName = data.ranks.overall.name;
    const level = levelName.includes("dan") ? "dan" : levelName.split("")[0];
    let cardTemplate = fs.readFileSync(
      path.join(__dirname + "../../templates/codewarscard.svg"),
      "utf8"
    );
    if (req.query.top_languages) {
      cardTemplate = setIcons(cardTemplate, data.ranks.languages);
    }
    res.send(
      cardTemplate
        .replace("{name}", req.query.name ? data.name : data.username)
        .replace("{rankName}", levelName)
        .replace("{clan}", data.clan)
        .replace("{leaderboardPosition}", data.leaderboardPosition)
        .replace("{honor}", data.honor)
        .replace("{score}", data.ranks.overall.score)
        .replace("{totalCompleted}", data.codeChallenges.totalCompleted)
        .replace(/{rankColor}/g, LevelColors[level])
    );
  } catch (err) {
    next(err);
  }
});

function setIcons(template, languages) {
  try {
    let x = -108;
    let icons_str = "";
    Object.keys(languages)
      .map((k) => {
        return { score: languages[k].score, name: k };
      })
      .sort((a, b) => a.score - b.score)
      .reverse()
      .slice(0, 3)
      .forEach((icon, i) => {
        let ic = "";
        try {
          if (specialIcons[icon.name]) {
            ic = sIcons
              .Get(specialIcons[icon.name].name)
              .svg.replace(specialIcons[icon.name].display_name, icon.name);
          } else {
            ic = sIcons.Get(icon.name).svg;
          }
        } catch (err) {
          ic = no_icon_found_template(icon.name);
          logger.error(`No Icon found for: ${icon.name}`);
        }
        if (i > 0) x += 60;
        icons_str =
          icons_str +
          icon_template(x)
            .replace("{svg}", ic)
            .replace(
              'viewBox="0 0 24 24"',
              'viewBox="0 0 150 150" class="icons" fill="#6795DE"'
            );
      });
    return setHeight(template).replace(
      "{icons}",
      top_languages_template(icons_str)
    );
  } catch (err) {
    logger.error(`Icons error: ${err}`);
    return template;
  }
}

function setHeight(template) {
  let to_replace_str = '<svg width="500" height="195" viewBox="0 0 500 195"';
  let h = (h) => `<svg width="500" height="${h}" viewBox="0 0 500 ${h}"`;
  return template.replace(to_replace_str, h(280));
}

export default router;