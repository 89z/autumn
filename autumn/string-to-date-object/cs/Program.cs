using System.Globalization;
using System;

class Program {
   static void Main() {
      CultureInfo enUS = new CultureInfo("en-US");
      DateTime dt2 = DateTime.ParseExact("2011-05-03 17:00:00", "yyyy-MM-dd HH:mm:ss", enUS);
   }
}
