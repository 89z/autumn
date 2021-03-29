let s = 'south north';
{ // example 1
   let r = /o../g;
   let a = [...s.matchAll(r)];
   console.log(a[0][0] == 'out', a[1][0] == 'ort');
}
{ // example 2
   let r = /o(..)/g;
   let a = [...s.matchAll(r)];
   console.log(a[0][1] == 'ut', a[1][1] == 'rt');
}
