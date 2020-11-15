function human_readable(in_n) {
   let out_n = 0;
   while (in_n > 1024) {
      in_n /= 1024;
      out_n++;
   }
   return in_n.toFixed(3) + ' ' + ' KMG'[out_n];
}

let s = human_readable(1264);
console.log(s == '1.234 K');
