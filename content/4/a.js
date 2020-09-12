let m = {year: 2019, month: 12};

console.log('example 1');
for (let s in m) {
   let n = m[s];
   console.log(s, n);
}

console.log('example 2');
for (let [s, n] of Object.entries(m)) {
   console.log(s, n);
}
