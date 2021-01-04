function sinceHours(s) {
   let n = Date.now() - Date.parse(s);
   return n / 36e5;
}

let n = sinceHours('2020-12-31');
console.log(n);
