void main() {
   var a = ['May', 'June'];
   // example 1
   var s1 = a.reduce((s, s1) => s + s1);
   // example 2
   var f = (String s, String s2) => s + s2;
   var s2 = a.reduce(f);
   // print
   print(s1 == 'MayJune' && s2 == 'MayJune');
}
