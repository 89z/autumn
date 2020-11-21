let o = Intl.NumberFormat('en', { notation: 'compact' });
// example 1
let s1 = o.format(1e3);
// example 2
let s2 = o.format(1e6);
// example 3
let s3 = o.format(1e9);
// example 4
let s4 = o.format(1e12);
// example 5
let s5 = o.format(1e15);
// print
console.log(s1 == '1K', s2 == '1M', s3 == '1B', s4 == '1T', s5 == '1000T');
