let s_uri = 'https://speedtest.lax.hivelocity.net';
fetch(s_uri).then(
   o_req => o_req.text()
).then(
   s_resp => console.log(s_resp)
);
