let a = ['May', 'June'];
// example A
let sA = a.reduce((sY, sZ) => sY + sZ);
// example B
let f = (sY, sZ) => sY + sZ;
let sB = a.reduce(f)
// print
console.log(sA == 'MayJune', sB == 'MayJune');
