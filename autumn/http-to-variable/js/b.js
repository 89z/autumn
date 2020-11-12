async function main() {
   let uri_s = 'https://speedtest.lax.hivelocity.net';
   let req_o = await fetch(uri_s);
   let resp_s = await req_o.text();
   console.log(resp_s);
}

main();
