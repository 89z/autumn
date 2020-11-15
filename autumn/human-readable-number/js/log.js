function human_readable(n) {
   let n2 = Math.log(n) / Math.log(1024);
   n2 = Math.floor(n2);
   return (n / 1024 ** n2).toFixed(3) + ['', ' k', ' M', ' G'][n2];
}

let s = human_readable(1264);
console.log(s == '1.234 k');
