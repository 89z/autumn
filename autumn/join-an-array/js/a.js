let a = ['west', 'east'];
{ // example 1
   let s = a.join(',');
   console.log(s == 'west,east');
}
{ // example 2
   let s = a.join();
   console.log(s == 'west,east');
}
