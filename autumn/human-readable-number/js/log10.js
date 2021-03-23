function numberFormat(d) {
   let e = Math.trunc(Math.log10(d) / 3);
   return (d / 1e3 ** e).toFixed(3) + ['', ' k', ' M', ' G'][e];
}

let s = numberFormat(9012345678);
console.log(s == '9.012 G');
