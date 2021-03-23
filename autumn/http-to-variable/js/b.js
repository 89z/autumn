async function main() {
   let uri = 'https://speedtest.lax.hivelocity.net';
   let req = await fetch(uri);
   let res = await req.text();
   console.log(res);
}

main();
