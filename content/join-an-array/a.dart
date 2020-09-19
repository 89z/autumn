void main() {
   var a = ['May', 'June'];
   // example 1
   var s1 = a.join(',');
   // example 2
   var s2 = a.join();
   // print
   print(s1 == 'May,June' && s2 == 'MayJune');
}
