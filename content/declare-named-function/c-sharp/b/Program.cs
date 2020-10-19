using System;

class Program {
   // example 1
   static bool f1(string s, char c) {
      return s[0] == c;
   }
   // example 2
   static bool f2(string s, char c) => s[0] == c;
   // print
   static void Main() {
      var b1 = f1("June", 'J');
      var b2 = f2("June", 'J');
      Console.WriteLine(b1 && b2);
   }
}
