'use strict';

function encode(m) {
   let p = new URLSearchParams(m);
   return p.toString();
}

let m = {month: 'March', day: 'Friday'};
let s = encode(m);
console.log(s === 'month=March&day=Friday');
