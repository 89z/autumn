let s = 'CgsMDQ=='
let t = atob(s);
console.log(t == '\x0a\x0b\x0c\x0d');
