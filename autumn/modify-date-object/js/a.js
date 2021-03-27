function addYear(d) {
   let n = d.getFullYear();
   d.setFullYear(n + 1);
}

let d = new Date;
addYear(d);
console.log(d);
