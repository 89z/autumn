let s = 'Sunday Monday';
{
   // example 1
   let o = s.matchAll(/..n/g);
   let a = Array.from(o);
   console.log(a[0][0] == 'Sun', a[1][0] == 'Mon');
}
{ // example 2
   let o = s.matchAll(/(..)n/g);
   let a = Array.from(o);
   console.log(a[0][1] == 'Su', a[1][1] == 'Mo');
}
