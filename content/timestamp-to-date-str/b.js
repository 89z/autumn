let o = new Date;
let s1 = o.toLocaleString(0, {weekday: 'short'});
let s2 = o.toLocaleString(0, {month: 'short', day: 'numeric'});
let s3 = o.toLocaleString(0, {year: 'numeric'});
console.log(s1, s2, s3);
