function numberFormat(n) {
   let n2 = Math.trunc(Math.log10(n) / 3);
   return (n / 1e3 ** n2).toFixed(3) + ['', ' k', ' M', ' G'][n2];
}

let s = numberFormat(9012345678);
console.log(s == '9.012 G');
