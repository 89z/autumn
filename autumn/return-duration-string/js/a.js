function format(d) {
   let mil = String(d % 1000).padStart(3, '0');
   d = Math.trunc(d / 1000);
   let sec = String(d % 60).padStart(2, '0');
   d = Math.trunc(d / 60);
   return String(d) + ' m ' + sec + ' s ' + mil + ' ms';
}

function main() {
   console.clear();
   let s = format(Date.now() - n);
   console.log(s);
}

let n = Date.now();
setInterval(main, 10);
