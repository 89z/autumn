let s = 'Sunday Monday';
{ // example 1
   let r = /..n/g;
   let a = [...s.matchAll(r)];
   console.log(a[0][0] == 'Sun', a[1][0] == 'Mon');
}
{ // example 2
   let r = /(..)n/g;
   let a = [...s.matchAll(r)];
   console.log(a[0][1] == 'Su', a[1][1] == 'Mo');
}
