using System;

class Program {
   static bool StartsWith(string s, char c) => s[0] == c;
   static void Main() {
      var b = StartsWith("June", 'J');
      Console.WriteLine(b);
   }
}
