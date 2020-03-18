let a1 = ['Sun', 'Mon'];
// example 1
for (let n1 = 0; n1 < a1.length; n1++) {
   console.log(a1[n1]);
}
// example 2
for (let n1 in a1) {
   console.log(a1[n1]);
}
// example 3
for (let s1 of a1) {
   console.log(s1);
}
// example 4
a1.forEach(s1 => console.log(s1));
