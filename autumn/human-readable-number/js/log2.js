function numberFormat(n) {
   let n2 = Math.log2(n) / 10;
   let n3 = Math.floor(n2);
   //return (n / 1000 ** n3).toFixed(3) + ['', ' k', ' M', ' G'][n3];
   return (n / 1000 ** n3) + ['', ' k', ' M', ' G'][n3];
}

let s = numberFormat(5678901234);
console.log(s);
