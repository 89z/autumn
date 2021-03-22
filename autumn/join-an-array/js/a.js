let a = ['May', 'June'];
{ // example 1
   let s = a.join(',');
   console.log(s == 'May,June');
}
{ // example 2
   let s = a.join();
   console.log(s == 'May,June');
}
