void main() {
   var a = ['May', 'June'];
   var f = (String s_acc, String s_cur) => s_acc + s_cur;
   // example 1
   var s = a.reduce(f);
   // example 2
   var s2 = '';
   for (var s_cur in a) {
      s2 = f(s2, s_cur);
   }
   // print
   print(s == 'MayJune' && s2 == 'MayJune');
}
