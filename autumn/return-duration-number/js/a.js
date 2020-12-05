function duration(new_o, old_o) {
   let new_n = new_o.getTime();
   let old_n = old_o.getTime();
   return (new_n - old_n) / 1000;
}

let o = new Date('2019-12-31');
let n = duration(new Date, o);
console.log(n);
