// example 1
class Jan {
   Sun = 10;
   Mon = 11;
}
let o1 = new Jan;
// example 2
let o2 = new class {
   Sun = 10;
   Mon = 11;
};
// print
console.log(o1, o2);
