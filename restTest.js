const autocannon = require('autocannon');
const fs = require('fs');

const payload = fs.readFileSync('./payload-small.json', 'utf8');

const instance = autocannon({
  url: 'http://localhost:8080/process',
  method: 'POST',
  connections: 100,
  duration: 20,
  headers: { 'Content-Type': 'application/json' },
  body: payload
});

autocannon.track(instance, { renderProgressBar: true });
