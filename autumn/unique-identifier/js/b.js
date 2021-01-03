function idDecode(s, year) {
   let t = new Date(year, 0);
   let x = t.getTime() + parseInt(s, 36) * 1000;
   return new Date(x);
}

let o = idDecode('itrzz', 2020);
console.log(o);
