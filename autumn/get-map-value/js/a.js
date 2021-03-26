let m = {month: 12, day: 31};

{ // example 1
   let n = m['day'];
   console.log(n == 31);
}
{ // example 2
   let n = m.day;
   console.log(n == 31);
}
