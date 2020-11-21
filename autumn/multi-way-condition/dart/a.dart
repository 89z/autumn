void main() {
   int n = 0x63 - 0x20;
   String s;

   switch (n) {
   case 0x41:
      s = 'A';
      break;
   case 0x42:
   case 0x62:
      s = 'B';
      break;
   default:
      s = String.fromCharCodes([n]);
   }

   print(s == 'C');
}
