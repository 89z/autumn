let s1 = '♠';
// bytes
let n1 = new TextEncoder().encode(s1).length;
// characters
let n2 = Array.from(s1).length;
// fails with SMP characters
let n3 = s1.length;
// print
console.log(n1, n2, n3);
