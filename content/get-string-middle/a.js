let s1 = 'Sunday';
let s2;
// example 1
s2 = s1.slice(1, 2);
console.log(s2 == 'u');
// example 2
s2 = s1.substring(1, 2);
console.log(s2 == 'u');
// example 3
s2 = s1.charAt(1);
console.log(s2 == 'u');
// example 4
s2 = s1[1];
console.log(s2 == 'u');
// example 5
s2 = s1.slice(1);
console.log(s2 == 'unday');
// example 6
s2 = s1.substring(1);
console.log(s2 == 'unday');
