void main() {
   var a = ['May', 'June'];
   // example 1
   var s = a.reduce((sa, sc) => sa + sc);
   // example 2
   var f = (String sa, String sc) => sa + sc;
   var s2 = a.reduce(f);
   // print
   print(s == 'MayJune' && s2 == 'MayJune');
}
