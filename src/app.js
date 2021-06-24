const express = require('express');
const morgan = require('morgan');
const helmet = require('helmet');
const axios = require('axios');
const cors = require('cors');
require('dotenv').config();
const middlewares = require('./middlewares');
const app = express();
const logger = require('./logger');

app.use(helmet());
app.use(morgan('tiny'));
app.use(cors());
app.use(express.json());

app.get('/', async(req, res) => {
    res.redirect('https://github.com/andreasvogt89/codewars_api');
});

/**
 * Rely to codewars user infos
 */
app.get('/codewars', async(req, res, next) => {
    const codewars_Url = 'https://www.codewars.com/api/v1/users/';
    try {
        const { data } = await axios.get(`${codewars_Url}/${req.query.user}`);
        res.setHeader("Content-Type", "image/svg+xml")
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
      font: 800 18px 'Segoe UI', Ubuntu, Sans-Serif; fill: #BB432C; 
      animation: scaleInAnimation 0.3s ease-in-out forwards;
      }
      
      .bold { font-weight: 700 }
      .icon {
      fill: #BB432C;
      display: none;
      }
      
      .rank-hex {
      stroke: #BB432C;
      fill: none;
      stroke-width: 2;
      opacity: 0.8;
      transform-origin: -10px 8px;
      transform: rotate(-90deg);
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
      transform="translate(400, 47.5)">
          <polygon class="rank-hex" points="300,150 225,280 75,280 0,150 75,20 225,20"></polygon>
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