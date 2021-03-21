let f = 7.5, i = 7;
{ // example 1
   let n = f / 2;
   console.log(n == 3.75);
}
{ // example 2
   let n = Math.trunc(f / 2);
   console.log(n == 3);
}
{ // example 3
   let n = Math.trunc(i / 2);
   console.log(n == 3);
}
{ // example 4
   let n = i / 2;
   console.log(n == 3.5);
}
