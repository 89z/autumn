let m = {month: 5, day: 4};
{ // example 1
   let n = m['day'];
   console.log(n == 4);
}
{ // example 2
   let n = m.day;
   console.log(n == 4);
}
