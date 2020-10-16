let n, s = 'minute';

switch (s) {
case 'hour':
   n = 23;
   break;
case 'minute':
case 'second':
   n = 59;
   break;
default:
   n = -1;
}

console.log(n == 59);
