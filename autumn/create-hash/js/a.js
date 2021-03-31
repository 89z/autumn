async function main() {
   let t = new TextEncoder;
   let a = t.encode('south north');
   let d = await crypto.subtle.digest('SHA-256', a);
   console.log(d);
}

main();
