using System;

class Time {
   int hours = 23;
   public int duration(int minutes) {
      return this.hours * 60 + minutes;
   }
}

class Program {
   static void Main() {
      var o = new Time();
      var n = o.duration(59);
      Console.WriteLine(n == 1439);
   }
}
