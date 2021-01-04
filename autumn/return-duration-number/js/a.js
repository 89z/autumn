function sinceHours(s) {
   let t = new Date(s);
   return (new Date - t) / 36e5;
}

let n = sinceHours('2020-12-31');
console.log(n);
