function human_readable(n3) {
   let n = Math.log2(n3) / 10;
   n = Math.floor(n);
   return (n3 / 1024 ** n).toFixed(3) + ['', ' k', ' M', ' G'][n];
}

let s = human_readable(1264);
console.log(s == '1.234 k');
