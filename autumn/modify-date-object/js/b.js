function addMonth(o) {
   let n = o.getMonth();
   o.setMonth(n + 1);
}

let o = new Date;
addMonth(o);
console.log(o);
