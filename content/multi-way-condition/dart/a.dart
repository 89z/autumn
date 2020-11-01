void main() {
   num n = 46 / 10;
   String s;

   switch (n) {
   case 6:
      s = 'six';
      break;
   case 5:
   case 4:
      s = 'five or four';
      break;
   default:
      s = n.toString();
   }

   print(s == '4.6');
}
