void main() {
   int n = 1 + 2;
   bool b;

   switch (n) {
   case 5:
      b = false;
      break;
   case 4:
   case 3:
      b = true;
      break;
   default:
      b = n < 3;
   }

   print(b);
}
