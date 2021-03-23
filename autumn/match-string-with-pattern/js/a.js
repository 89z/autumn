let s = 'January';
{ // example 1
   let b = /^J/.test(s);
   console.log(b);
}
{ // example 2
   let b = /ja/i.test(s);
   console.log(b);
}
