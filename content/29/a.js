let n = 366 * 24 * 60 * 60;
let o = new Date(n * 1000);
let s_d = o.toLocaleString(0, {weekday: 'short'});
let s_m = o.toLocaleString(0, {month: 'short', day: 'numeric'});
let s_y = o.toLocaleString(0, {year: 'numeric'});
let s = s_d + ' ' + s_m + ' ' + s_y;
console.log(s == 'Fri Jan 1 1971');
