void main() {
   num n = 46 / 10;
   String s;

   switch (n) {
   case 7:
      s = 'seven';
      break;
   case 6:
   case 5:
      s = 'six or five';
      break;
   default:
      s = n.toString();
   }

   print(s);
}
