function Day() {
   this.Sunday = 10;
}
let o = new Day;
// example 1
let b1 = 'Sunday' in o;
// example 2
let b2 = o.hasOwnProperty('Sunday');
// print
console.log(b1, b2);
