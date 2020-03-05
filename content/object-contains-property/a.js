class day {
   sun = 10;
   mon = 11;
}
let o1 = new day;
// example 1
let b1 = 'sun' in o1;
// example 2
let b2 = o1.hasOwnProperty('sun');
// print
console.log(b1, b2);
