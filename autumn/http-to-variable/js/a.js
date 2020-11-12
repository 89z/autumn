let uri_s = 'http://speedtest.lax.hivelocity.net';

fetch(uri_s).then(
   req_o => req_o.text()
).then(
   resp_s => console.log(resp_s)
);
