function duration(u, t) {
   let new_n = u.getTime();
   let old_n = t.getTime();
   return (new_n - old_n) / 1000;
}

let t = new Date('2020-12-31');
let n = duration(new Date, t);
console.log(n);
