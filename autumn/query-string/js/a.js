let m = {one: 'odd', two: 'even'};
let o = new URLSearchParams(m);
let s = o.toString();
console.log(s === 'one=odd&two=even');
