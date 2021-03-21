'use strict';
let s = 'Sunday Monday';
{ // example 1
   let a = s.match(/..n/);
   console.log(a[0] == 'Sun');
}
{ // example 2
   let a = s.match(/..n/g);
   console.log(a[0] == 'Sun', a[1] == 'Mon');
}
{ // example 3
   let a = s.match(/(..)n/);
   console.log(a[1] == 'Su');
}
