let f = Intl.NumberFormat('en', { notation: 'compact' });
{ // example 1
   let s = f.format(1e3);
   console.log(s == '1K');
}
{ // example 2
   let s = f.format(1e6);
   console.log(s == '1M');
}
{ // example 3
   let s = f.format(1e9);
   console.log(s == '1B');
}
{ // example 4
   let s = f.format(1e12);
   console.log(s == '1T');
}
{ // example 5
   let s = f.format(1e15);
   console.log(s == '1000T');
}
