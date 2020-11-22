void main() {
   String s = '\x43';
   int n;

   switch (s) {
   case 'A':
      n = 0x41;
      break;
   case 'B':
   case 'b':
      n = 0x42;
      break;
   default:
      n = s.codeUnitAt(0);
   }

   print(n == 0x43);
}
