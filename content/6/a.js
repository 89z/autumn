let o1 = new Date('2019-12-31T00:00:00');
let o2 = new Date('2019-12-31T23:59:59');
let o3 = new Date(o2 - o1);
let s_h = o3.getUTCHours().toString();
let s_m = o3.getUTCMinutes().toString();
let s_s = o3.getUTCSeconds().toString();
let s = s_h + ':' + s_m + ':' + s_s;
console.log(s);
