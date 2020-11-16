function numberFormat(n) {
   let n2 = 0;
   while (n > 1024) {
      n /= 1024;
      n2++;
   }
   return n.toFixed(3) + ['', ' k', ' M', ' G'][n2];
}

let s = numberFormat(1264);
console.log(s == '1.234 k');
