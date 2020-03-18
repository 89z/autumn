class Day {
   Sun = 10;
   Mon = 11;
}
let o1 = new Day;
// example 1
let b1 = 'Sun' in o1;
// example 2
let b2 = o1.hasOwnProperty('Sun');
// print
console.log(b1, b2);
