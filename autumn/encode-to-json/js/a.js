let a = ['/', '📗'];
{ // example 1
   let s = JSON.stringify(a);
   console.log(s);
}
{ // example 2
   let s = JSON.stringify(a, null, 1);
   console.log(s);
}
