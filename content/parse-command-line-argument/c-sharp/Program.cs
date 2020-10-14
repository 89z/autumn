using System;

class Program {
   static void Main(string begin, string end, string input) {
      if (input == "") {
         Console.WriteLine(@"slice [flags] <string>
-len int
   length (default 1)
-start int
   index");
         Environment.Exit(1);
      }
      if (len == 0) {
         len = 1;
      }
      var s_out = s_in.Substring(start, len);
      Console.WriteLine(s_out);
   }
}
