let s = 'south north';
{ // example 1
   let a = s.match(/o../);
   console.log(a[0] == 'out');
}
{ // example 2
   let a = s.match(/o../g);
   console.log(a[0] == 'out', a[1] == 'ort');
}
{ // example 3
   let a = s.match(/o(..)/);
   console.log(a[1] == 'ut');
}
