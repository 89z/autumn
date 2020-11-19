'use strict';
let s = 'Sunday Monday';

// example 1
let a1 = s.match(/..n/);
console.log(a1[0] == 'Sun');

// example 2
let a2 = s.match(/..n/g);
console.log(a2[0] == 'Sun', a2[1] == 'Mon');

// example 3
let a3 = s.match(/(..)n/);
console.log(a3[1] == 'Su');
