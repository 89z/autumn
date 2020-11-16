function numberFormat(n) {
   let n2 = Math.log(n) / Math.log(1024);
   let n3 = Math.floor(n2);
   return (n / 1024 ** n3).toFixed(3) + ['', ' k', ' M', ' G'][n3];
}

let s = numberFormat(1264);
console.log(s == '1.234 k');
