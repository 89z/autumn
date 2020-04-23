let s1 = 'http://speedtest.lax.hivelocity.net';
// example 1
fetch(s1).
then(o1 => o1.text()).
then(s2 => console.log(s2));
// example 2
(async () => {
   let o1 = await fetch(s1);
   let s2 = await o1.text();
   console.log(s2);
})();
