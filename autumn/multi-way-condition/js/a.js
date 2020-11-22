let s = '\x43', n;

switch (s) {
case 'A':
   n = 0x41;
   break;
case 'B':
case 'b':
   n = 0x42;
   break;
default:
   n = s.charCodeAt();
}

console.log(n == 0x43);
