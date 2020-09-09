let m = {year: 2020, month: 9};

console.log('example 1');
for (let s in m) {
   console.log(s);
}

console.log('example 2');
for (let [s, n] of Object.entries(m)) {
   console.log(s, n);
}
