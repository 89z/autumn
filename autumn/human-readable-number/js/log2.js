function human_readable(n) {
   let n2 = Math.log2(n) / 10;
   n2 = Math.floor(n2);
   return (n / 1024 ** n2).toFixed(3) + ' ' + ' KMG'[n2];
}

let s = human_readable(1264);
console.log(s == '1.234 K');
