'use strict';

function format(n) {
   let mil_s = String(n % 1000).padStart(3, '0');
   n = Math.trunc(n / 1000);
   let sec_s = String(n % 60).padStart(2, '0');
   n = Math.trunc(n / 60);
   return String(n) + ' m ' + sec_s + ' s ' + mil_s + ' ms';
}

function main() {
   console.clear();
   let s = format(Date.now() - n);
   console.log(s);
}

let n = Date.now();
setInterval(main, 10);
