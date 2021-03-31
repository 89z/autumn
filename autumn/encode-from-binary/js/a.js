let a = [0xA, 0xB, 0xC, 0xD];
let s = a.map(n => n.toString(16).padStart(2, '0')).join('');
console.log(s == '0a0b0c0d');
