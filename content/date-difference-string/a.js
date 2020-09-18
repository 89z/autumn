let o1 = new Date('2019-12-31T00:00:00');
let o1a = new Date('2019-12-31T23:59:59');
let o = new Date(o1a - o1);
let s_h = o.getUTCHours().toString();
let s_m = o.getUTCMinutes().toString();
let s_s = o.getUTCSeconds().toString();
let s = s_h + ':' + s_m + ':' + s_s;
console.log(s);
