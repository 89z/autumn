'use strict';

function idEncode(n) {
   let t = new Date(n, 0);
   return Math.trunc(
      (new Date - t) / 1000
   ).toString(36);
}

let s = idEncode(2021);
console.log(s);
