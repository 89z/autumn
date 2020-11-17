let o = Intl.NumberFormat('en', { notation: 'compact' });
// example 1
let s1 = o.format(9012);
// example 2
let s2 = o.format(9012345);
// example 3
let s3 = o.format(9012345678);
// print
console.log(s1 == '9K', s2 == '9M', s3 == '9B');
