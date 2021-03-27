function addMonth(d) {
   let n = d.getMonth();
   d.setMonth(n + 1);
}

let d = new Date;
addMonth(d);
console.log(d);
