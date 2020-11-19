function f(r) {
   let o = 'Sunday Monday'.matchAll(r);
   return Array.from(o);
}

// example 1
let a1 = f(/..n/g);
console.log(a1[0][0] == 'Sun', a1[1][0] == 'Mon');

// example 2
let a2 = f(/(..)n/g);
console.log(a2[0][1] == 'Su', a2[1][1] == 'Mo');
