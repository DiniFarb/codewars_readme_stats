const express = require('express');
const helmet = require('helmet');
const axios = require('axios');
const cors = require('cors');
require('dotenv').config();
const middlewares = require('./middlewares');
const app = express();
const logger = require('./logger');
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
app.use(helmet());
app.use(cors());

app.get('/', async(req, res) => {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    logger.info(`redirect: ${ip}`);
    res.redirect('https://github.com/andreasvogt89/codewars_api');
});

app.get('/codewars', async(req, res, next) => {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    logger.info(`crad request: ${ip}`);
    const codewars_Url = 'https://www.codewars.com/api/v1/users/';
    try {
        const { data } = await axios.get(`${codewars_Url}/${req.query.user}`);
        res.setHeader("Content-Type", "image/svg+xml");
        const level = data.ranks.overall.name.split('')[0];
        res.send(`<svg width="495" height="195" viewBox="0 0 495 195" fill="none" xmlns="http://www.w3.org/2000/svg">
        <style>
          .header {
            font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif;
            fill: #F1F5F3;
            animation: fadeInAnimation 0.8s ease-in-out forwards;
          }
          
      .stat {
      font: 600 14px 'Segoe UI', Ubuntu, "Helvetica Neue", Sans-Serif; fill: #BB432C;
      }
      .stats {
      opacity: 0;
      animation: fadeInAnimation 0.3s ease-in-out forwards;
      }
      .rank-text {
      fill: ${LevelColors[level]};
      font: 600 30px 'Segoe UI', Ubuntu, Sans-Serif; 
      animation: scaleInAnimation 0.3s ease-in-out forwards;
      }
      
      .bold { font-weight: 700 }
      .icon {
      fill: #BB432C;
      display: none;
      }
      
      .rank-hex {
      stroke: ${LevelColors[level]};
      fill: #181919;
      stroke-width: 3;
      opacity: 0.8;
      animation: rankAnimation 1s forwards ease-in-out;
      }
      
      @keyframes rankAnimation {
      from {
        stroke-dashoffset: 251.32741228718345;
      }
      to {
        stroke-dashoffset: 123.09285659688518;
      }
      }
      
      
      
          
      /* Animations */
      @keyframes scaleInAnimation {
      from {
        transform: translate(-5px, 5px) scale(0);
      }
      to {
        transform: translate(-5px, 5px) scale(1);
      }
      }
      @keyframes fadeInAnimation {
      from {
        opacity: 0;
      }
      to {
        opacity: 1;
      }
      }
      
          
        </style>
      
        undefined
      
        <rect
          data-testid="card-bg"
          x="0.5"
          y="0.5"
          rx="4.5"
          height="99%"
          stroke="#BB432C"
          width="494"
          fill="#262729"
          stroke-opacity="0"
        />
      
        
      <g
        transform="translate(25, 35)"
      >
        <g transform="translate(0, 0)">
      <text
        x="0"
        y="0"
        class="header"
      >${data.name}'s Codewars Stats</text>
      </g>
      </g>
      
      
        <g
          transform="translate(0, 55)"
        >
          
      <g 
      transform="translate(400, 50)">
        <polygon class="rank-hex" points="-60,7.5 -45,-20 35,-20 50,7.5 35,35 -45,35"></polygon>
        <g class="rank-text">
          <text
            x="0"
            y="0"
            alignment-baseline="central"
            dominant-baseline="central"
            text-anchor="middle"
          >
            ${data.ranks.overall.name}
          </text>
        </g>
      </g>
      
      <svg x="0" y="0">
      <g transform="translate(0, 0)">
      <g class="stats" style="animation-delay: 450ms" transform="translate(25, 0)">
      
      <text class="stat bold"  y="12.5">Clan:</text>
      <text 
        class="stat" 
        x="170" 
        y="12.5" 
      >${data.clan}</text>
      </g>
      </g><g transform="translate(0, 25)">
      <g class="stats" style="animation-delay: 600ms" transform="translate(25, 0)">
      
      <text class="stat bold"  y="12.5">Leader board position:</text>
      <text 
        class="stat" 
        x="170" 
        y="12.5" 
      >${data.leaderboardPosition}</text>
      </g>
      </g><g transform="translate(0, 50)">
      <g class="stats" style="animation-delay: 750ms" transform="translate(25, 0)">
      
      <text class="stat bold"  y="12.5">Honor:</text>
      <text 
        class="stat" 
        x="170" 
        y="12.5" 
      >${data.honor}</text>
      </g>
      </g><g transform="translate(0, 75)">
      <g class="stats" style="animation-delay: 900ms" transform="translate(25, 0)">
      
      <text class="stat bold"  y="12.5">Score:</text>
      <text 
        class="stat" 
        x="170" 
        y="12.5" 
      >${data.ranks.overall.score}</text>
      </g>
      </g><g transform="translate(0, 100)">
      <g class="stats" style="animation-delay: 1050ms" transform="translate(25, 0)">
      
      <text class="stat bold"  y="12.5">Solved Katas:</text>
      <text 
        class="stat" 
        x="170" 
        y="12.5" 
      >${data.codeChallenges.totalCompleted}</text>
      </g>
      </g>
      </svg>

</g>
</svg>`)
    } catch (err) {
        next(err);
    }
});

app.use(middlewares.notFound);
app.use(middlewares.errorHandler);

module.exports = app;