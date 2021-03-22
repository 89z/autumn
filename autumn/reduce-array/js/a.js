let a = ['May', 'June'];
{ // example 1
   let f = (s, t) => s + t;
   let s = a.reduce(f);
   console.log(s == 'MayJune');
}
{ // example 2
   let f = (n, s) => n + s.length;
   let n = a.reduce(f, 0);
   console.log(n == 7);
}
