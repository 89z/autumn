'use strict';

function format(n) {
   let mil = String(n % 1000).padStart(3, '0');
   n = Math.trunc(n / 1000);
   let sec = String(n % 60).padStart(2, '0');
   n = Math.trunc(n / 60);
   return String(n) + ' m ' + sec + ' s ' + mil + ' ms';
}

function main() {
   console.clear();
   let s = format(Date.now() - n);
   console.log(s);
}

let n = Date.now();
setInterval(main, 10);
