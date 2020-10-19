using System;

static class Program {
   // example 1
   static bool f1(this string s, char c) {
      return s[0] == c;
   }
   // example 2
   static bool f2(this string s, char c) => s[0] == c;
   // print
   static void Main() {
      var b1 = "June".f1('J');
      var b2 = "June".f2('J');
      Console.WriteLine(b1 && b2);
   }
}
