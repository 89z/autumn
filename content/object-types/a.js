// example 1
class jan {
   sun = 10;
   mon = 11;
}
let o1 = new jan;
// example 2
let o2 = new class {
   sun = 10;
   mon = 11;
};
// print
console.log(o1, o2);
