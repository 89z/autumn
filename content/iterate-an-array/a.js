let a = ['Sunday', 'Monday'];
// example 1
for (let n = 0; n < a.length; n++) {
   console.log(a[n]);
}
// example 2
for (let n in a) {
   console.log(a[n]);
}
// example 3
for (let s of a) {
   console.log(s);
}
// example 4
a.forEach(s => console.log(s));
