/* eslint-disable no-undef */
import request from 'supertest';
import app from '../src/app.js';
import expect from 'expect.js';

describe('GET /', () => {
  it('BASE is redirected',(done) => {
    request(app)
      .get('/')
      .redirects(1)
      .expect(301, done);
  });
});

describe('GET with unknown route', function() {
  it('responds with 404 error message', function(done) {
    request(app)
      .get('/free')
      .expect(404)      
      .expect((res)=>{
        expect(res.body).to.be('Not Found - /free');
      })
      .end(done);
  });
});

describe('GET /codewars without user params', function() {
  it('responds with 500 and error message', function(done) {
    request(app)
      .get('/codewars')
      .expect(500)
      .expect((res)=>{
        expect(res.body).to.be('Missing Query param => [user={yourname}]');
      })
      .end(done);
  });
});

describe('GET /codewars with unknown user', function() {
  it('responds with 500 and error message', function(done) {
    request(app)
      .get('/codewars?user=dnsvjheuowflaschajaf')
      .expect(500)
      .expect((res)=>{
        expect(res.body).to.be('codewars API failure: Error: Request failed with status code 404');
      })
      .end(done);
  });
});

describe('GET /codewars with known user', function() {
  it('responds with 200 and svg', function(done) {
    request(app)
      .get('/codewars?user=andreasvogt89')
      .expect('Content-Type', 'image/svg+xml; charset=utf-8')
      .expect(200,done);
  });
});
