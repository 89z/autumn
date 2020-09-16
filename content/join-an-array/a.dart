void main() {
   var a = ['May', 'June'];
   // example 1
   var s = a.join(',');
   // example 2
   var s2 = a.join();
   // print
   print(s == 'May,June' && s2 == 'MayJune');
}
