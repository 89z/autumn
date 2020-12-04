using C = System.Console;
using G = System.Collections.Generic;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      long tenM = 10 * 1000 * 1000;
      long sec = dt.ToFileTimeUtc() / tenM;
      long epoch = sec - 11644473600;
      long ft = (epoch + 11644473600) * tenM;
      DateTime dt2 = DateTime.FromFileTimeUtc(ft);
   }
}
