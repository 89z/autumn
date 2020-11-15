function human_readable(n3) {
   let n = Math.log10(n3) / 3;
   n = Math.floor(n);
   return (n3 / 1000 ** n).toFixed(3) + ['', ' k', ' M', ' G'][n];
}

let s = human_readable(1264);
console.log(s == '1.264 k');
