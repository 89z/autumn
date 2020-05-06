let o1 = new Date;
// example 1
let s1 = o1.toISOString();
console.log(s1);
// example 2
let s2 = o1.toDateString();
console.log(s2);
// example 3
let s3 = o1.toLocaleString(0, {weekday: 'short'});
let s4 = o1.toLocaleString(0, {month: 'short', day: 'numeric'});
let s5 = o1.toLocaleString(0, {year: 'numeric'});
console.log(s3, s4, s5);
