let s, n = 10;

switch (n) {
case 12:
   s = 'all';
   break;
case 11:
case 10:
   s = 'some';
   break;
default:
   s = 'none';
}

console.log(s == 'some');
