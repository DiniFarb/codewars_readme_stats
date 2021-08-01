const express = require('express');
const logger = require('../logger');
const router = express.Router();
const axios = require('axios');
const fs = require('fs');
const path = require('path');
const simpleIcons = require('simple-icons');

const LevelColors = {
    1: "#866CC7",
    2: "#866CC7",
    3: "#3C7EBB",
    4: "#3C7EBB",
    5: "#ECB613",
    6: "#ECB613",
    7: "#E6E6E6",
    8: "#E6E6E6"
};

let top_languages_template = (icons) => `
<g transform="translate(130, 190)">
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
</g>`

router.get('/', async(req, res, next) => {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    logger.info(`crad request: ${ip}`);
    const codewars_Url = 'https://www.codewars.com/api/v1/users';
    try {
        const { data } = await axios.get(`${codewars_Url}/${req.query.user}`);
        res.setHeader("Content-Type", "image/svg+xml");
        res.setHeader("Cache-Control", `public, max-age=7200`);
        const level = data.ranks.overall.name.split('')[0];
        let cardTemplate = fs.readFileSync(path.join(__dirname + '../../templates/codewarscard.svg'), 'utf8');
        //TODO: This does not work yet
        if(req.query.top_languages){
          cardTemplate = setIcons(cardTemplate,data.ranks.languages);
        }
        res.send(cardTemplate
          .replace('{name}',req.query.name ? data.name : data.username)
          .replace('{rankName}',data.ranks.overall.name)
          .replace('{clan}',data.clan)
          .replace('{leaderboardPosition}',data.leaderboardPosition)
          .replace('{honor}',data.honor)
          .replace('{score}',data.ranks.overall.score)
          .replace('{totalCompleted}',data.codeChallenges.totalCompleted)
          .replace(/{rankColor}/g,LevelColors[level])
        );
    } catch (err) {
        next(err);
    }
});

function setIcons(template,languages){
  try {
    let icons = Object.keys(languages);
    let icons_str = "";
    let temp = (x) => `
    <g transform="translate(${x},20)" >
    {svg}
    </g>`;
    let x = -100;
    icons.forEach((icon, i)=>{
      let ic = `<svg viewBox="0 0 24 24"><text>${icon}</text></svg>`;
      try {
        ic = simpleIcons.Get(icon).svg;
      } catch (err){
        logger.error(`No Icon found for: ${icon}`);
      }
      if(i > 0) x += 60;
      let template = temp(x);
      icons_str = icons_str + template
      .replace('{svg}',ic)
      .replace('viewBox="0 0 24 24"','viewBox="0 0 150 150" class="icons" fill="#6795DE"');
    });
    return setHeight(template).replace('{icons}', top_languages_template(icons_str));
  } catch (err){
    logger.error(`Icons error: ${err}`);
    return template
  }
}

function setHeight(template){
  let to_replace_str = '<svg width="500" height="195" viewBox="0 0 500 195"'
  let h = (h) => `<svg width="500" height="${h}" viewBox="0 0 500 ${h}"`;
  return template.replace(to_replace_str,h(280));
}

module.exports = router;