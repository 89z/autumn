function numberFormat(n) {
   let n2 = 0;
   while (n > 1000) {
      n /= 1000;
      n2++;
   }
   return n.toFixed(3) + ['', ' k', ' M', ' G'][n2];
}

let s = numberFormat(9012345678);
console.log(s == '9.012 G');
