using System;

class Program {
   // example 1
   static Func<string, char, bool> f1 = (s, c) => {
      return s[0] == c;
   };
   // example 2
   static Func<string, char, bool> f2 = (s, c) => s[0] == c;
   // print
   static void Main() {
      var b1 = f1("June", 'J');
      var b2 = f2("June", 'J');
      Console.WriteLine(b1 && b2);
   }
}
