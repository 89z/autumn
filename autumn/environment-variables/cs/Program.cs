using System;

class Program {
   static void Main() {
      var s = "USERPROFILE";
      var s1 = Environment.GetEnvironmentVariable(s);
      Console.WriteLine(s1 == @"C:\Users\Steven");
   }
}
