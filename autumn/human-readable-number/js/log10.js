function human_readable(n) {
   let n2 = Math.log10(n) / 3;
   n2 = Math.floor(n2);
   return (n / 1000 ** n2).toFixed(3) + ' ' + ' KMG'[n2];
}

let s = human_readable(1264);
console.log(s == '1.264 K');
