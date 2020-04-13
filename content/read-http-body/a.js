let s1 = 'http://speedtest.lax.hivelocity.net';
// example 1
fetch(s1).
then(r1 => r1.text()).
then(s2 => console.log(s2));
// example 2
(async () => {
   let r1 = await fetch(s1);
   let s2 = await r1.text();
   console.log(s2);
})();
