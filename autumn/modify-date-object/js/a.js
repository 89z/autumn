function addYear(o) {
   let n = o.getFullYear();
   o.setFullYear(n + 1);
}

let o = new Date;
addYear(o);
console.log(o);
