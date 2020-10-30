// example 1
function f(n) {
   return n > 0 ? '+' : '-';
}
console.log(f(1) == '+', f(-1) == '-');
// example 2
function f(n) {
   return n >= 0 ? n <= 0 ? 'zero' : '+' : '-';
}
console.log(f(1) == '+', f(0) == 'zero', f(-1) == '-');
