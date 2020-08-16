async function main() {
   let s_uri = 'https://speedtest.lax.hivelocity.net';
   let o_req = await fetch(s_uri);
   let s_resp = await o_req.text();
   console.log(s_resp);
}
main();
