let c1 = new Set(['sun', 'mon']);
let c2 = {'sun': true, 'mon': true};
// example 1
c1.forEach(s1 => console.log(s1));
// example 2
for (let s1 of c1) {
   console.log(s1);
}
// example 3
for (let s1 in c2) {
   console.log(s1);
}
