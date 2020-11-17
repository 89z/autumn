function numberFormat(n) {
   let n2 = Math.log(n) / Math.log(1000);
   let n3 = Math.floor(n2);
   return (n / 1000 ** n3).toFixed(3) + ['', ' k', ' M', ' G'][n3];
}

let s = numberFormat(1234);
console.log(s);
