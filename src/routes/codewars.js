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

const specialIcons = {
  shell: {
    name:"windowsterminal",
    display_name:"Windows Terminal"
  }
}

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
</g>`

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
    let x = -108;
    let icons_str = "";
    Object.keys(languages)
    .map(k=>{return {score:languages[k].score,name:k}})
    .sort((a,b)=> a.score + b.score)
    .slice(0,3)
    .forEach((icon, i)=>{
      let ic = ""
      try {
        if(specialIcons[icon.name]){
          ic = simpleIcons.Get(specialIcons[icon.name].name)
          .svg.replace(specialIcons[icon.name].display_name,icon.name);
        } else {
          ic = simpleIcons.Get(icon.name).svg;
        }
      } catch (err){
        ic = no_icon_found_template(icon.name);
        logger.error(`No Icon found for: ${icon.name}`);
      }
      if(i > 0) x += 60;
      icons_str = icons_str + 
      icon_template(x)
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