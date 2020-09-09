main() {
   var s, n = 1;
   switch (n) {
   case 3:
      s = 'all';
      break;
   case 2:
   case 1:
      s = 'some';
      break;
   default:
      s = 'none';
   }
   print(s == 'some');
}
