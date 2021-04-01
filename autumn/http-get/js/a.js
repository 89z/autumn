let uri = 'http://speedtest.lax.hivelocity.net';

fetch(uri).then(
   req => req.text()
).then(
   res => console.log(res)
);
