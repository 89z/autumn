let a1, s2, s1 = 'Wednesday';
// example 1
a1 = s1.match(/e./);
s2 = JSON.stringify(a1);
console.log(s2 == '["ed"]');
// example 2
a1 = s1.match(/e./g);
console.log(a1 == ['ed', 'es']);
// example 3
o1 = s1.matchAll(/e./g);
console.log(a1 == [['ed'], ['es']]);
// example 4
o1 = s1.match(/e(..)/);
console.log(a1 == ['edn', 'dn']);
// example 5
o1 = s1.matchAll(/e(..)/g);
console.log(a1 == [['edn', 'dn'], ['esd', 'sd']]);
