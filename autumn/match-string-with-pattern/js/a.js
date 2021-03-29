let s = 'South North';
{ // example 1
   let b = /^S/.test(s);
   console.log(b);
}
{ // example 2
   let b = /so/i.test(s);
   console.log(b);
}
