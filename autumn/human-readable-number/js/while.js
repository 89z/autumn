function numberFormat(n) {
   let [n2, n3] = [n, 0];
   while (n2 >= 1000) {
      n2 /= 1000;
      n3++;
   }
   return n2.toFixed(3) + ['', ' k', ' M', ' G'][n3];
}

let s = numberFormat(9012345678);
console.log(s == '9.012 G');
