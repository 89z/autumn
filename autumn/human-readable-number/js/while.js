function numberFormat(d) {
   let [e, f] = [d, 0];
   while (e >= 1000) {
      e /= 1000;
      f++;
   }
   return e.toFixed(3) + ['', ' k', ' M', ' G'][f];
}

let s = numberFormat(9012345678);
console.log(s == '9.012 G');
