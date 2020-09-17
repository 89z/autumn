function f(sa, sc) {
   return sa + sc;
}

let a = ['May', 'June'];
let s = '';

for (let sc of a) {
   s = f(s, sc);
}

console.log(s == 'MayJune');
