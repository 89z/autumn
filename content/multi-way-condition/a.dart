main() {
   var s = 'Sun';
   switch (s) {
   case 'Fri':
      print(1);
      break;
   case 'Sat':
   case 'Sun':
      print(2);
      break;
   default:
      print(0);
   }
}
