let s1 = 'Sunday';
let s2;
// example 1
s2 = s1.slice(0, 1);
console.log(s2 == 'S');
// example 2
s2 = s1.substring(0, 1);
console.log(s2 == 'S');
// example 3
s2 = s1.charAt(0);
console.log(s2 == 'S');
// example 4
s2 = s1.charAt();
console.log(s2 == 'S');
// example 5
s2 = s1[0];
console.log(s2 == 'S');
// example 6
s2 = s1.slice(0, 2);
console.log(s2 == 'Su');
// example 7
s2 = s1.substring(0, 2);
console.log(s2 == 'Su');
