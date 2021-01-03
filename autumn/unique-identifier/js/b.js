function idEncode(year) {
   let t = new Date(year, 0);
   let x = (new Date - t) / 1000;
   return Math.trunc(x).toString(36);
}

let s = idEncode(2020);
console.log(s);
