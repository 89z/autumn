// example 1
let a1 = ['Sun', 'Mon'];
let t1 = new Set(a1);
for (let s1 of t1) {
   console.log(s1);
}
// example 2
let t2 = {'Sun': true, 'Mon': true};
for (let s1 in t2) {
   console.log(s1);
}
// example 3
let t3 = {'Sun': true, 'Mon': true};
t2.forEach(s1 => console.log(s1));
