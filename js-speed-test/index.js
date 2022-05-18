const http = require('http');

const server = http.createServer((req, res) => {
  const incl = new URL(`http://localhost:4000${req.url}`).searchParams.get('include') === 'true';
  const to = +new URL(`http://localhost:4000${req.url}`).searchParams.get('to') || 10000000;

  const start = Date.now();
  const ps = speedTest(to);
  const elapsed = Date.now() - start;

  const response = {
    numbersTested: {
      start: 0,
      to: to,
    },
    elapsedTime: `${elapsed}ms`,
    amount: ps.length,
  };

  if (incl) response.results = ps;

  res.statusCode = 200;
  res.setHeader('content-type', 'application/json');
  res.setHeader('access-control-allow-credentials', 'true');
  res.setHeader('access-control-allow-headers', '*');
  res.setHeader('access-control-allow-methods', '*');
  res.setHeader('access-control-allow-origin', '*');
  res.setHeader('cache-control', 'no-cache');
  res.setHeader('content-type', 'application/json');
  res.end(JSON.stringify(response));
});

server.listen(4000, () => console.log(`Server running at port ${4000}`));

function speedTest(to) {
  const p = [];

  for (let i = 0; i < to; i++) {
    if (prime(i)) p.push(i);
  }
  return p;
}

function prime(n) {
  const sqrt = Math.sqrt(n);
  for (let i = 2; i <= sqrt; i++) {
    if (n % i == 0) return false;
  }
  return true;
}
