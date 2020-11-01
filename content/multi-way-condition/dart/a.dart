void main() {
   int n = 1 + 2;
   String s;

   switch (n) {
   case 5:
      s = 'five';
      break;
   case 4:
   case 3:
      s = 'four or three';
      break;
   default:
      s = n.toString();
   }

   print(s == 'four or three');
}
