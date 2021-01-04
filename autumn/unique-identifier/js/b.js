function idDecode(s, year) {
   let d = Date.UTC(year) + parseInt(s, 36) * 1000;
   return new Date(d);
}

let o = idDecode('6dv3d', 2020);
console.log(o);
