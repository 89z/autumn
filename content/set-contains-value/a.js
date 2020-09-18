// example A
let tA = new Set(['May', 'June']);
let bA = tA.has('June');
// example B
let tB = {May: true, June: true};
let bB = 'June' in tB;
// print
console.log(bA, bB);
