// example 1
let c1 = {'Sun': true, 'Mon': true};
for (let s1 in c1) {
   console.log(s1);
}
// example 2
let c2 = {'Sun': true, 'Mon': true};
c2.forEach(s1 => console.log(s1));
// example 3
let c3 = new Set(['Sun', 'Mon']);
for (let s1 of c3) {
   console.log(s1);
}
