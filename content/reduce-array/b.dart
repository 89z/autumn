String f(String sa, String sc) {
   return sa + sc;
}

void main() {
   var a = ['May', 'June'];
   var s = '';
   for (var sc in a) {
      s = f(s, sc);
   }
   print(s == 'MayJune');
}
