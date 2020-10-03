let a = ['May', 'June'];
// example 1
for (let n = 0; n < a.length; n++) {
   let s = a[n];
   console.log(n, s);
}
// example 2
for (let n in a) {
   let s = a[n];
   console.log(n, s);
}
// example 3
for (let s of a) {
   console.log(s);
}
// example 4
let f = s => console.log(s);
a.forEach(f);
