using System;

class Radix36 {
   string digit = "0123456789abcdefghijklmnopqrstuvwxyz";

   public string encode(int input) {
      string output = "";
      do {
         output = this.digit[input % 36] + output;
         input /= 36;
      } while (input > 0);
      return output;
   }

   public int decode(string input) {
      int output = 0;
      foreach (var letter in input) {
         output = output * 36 + this.digit.IndexOf(letter);
      }
      return output;
   }
}

class Program {
   static void Main() {
      var o = new Radix36();
      var s = o.encode(1577858399);
      var n = o.decode("q3ezbz");
      Console.WriteLine(s == "q3ezbz" && n == 1577858399);
   }
}
