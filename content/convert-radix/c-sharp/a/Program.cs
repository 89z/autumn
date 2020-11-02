using System;

class Radix {
   string sDigit = "0123456789abcdefghijklmnopqrstuvwxyz";

   public string encode(int nIn, int nBase) {
      string sOut = "";
      do {
         sOut = this.sDigit[nIn % nBase] + sOut;
         nIn /= nBase;
      } while (nIn > 0);
      return sOut;
   }

   public int decode(string sIn, int nBase) {
      int nOut = 0;
      foreach (char cValue in sIn) {
         nOut = nOut * nBase + this.sDigit.IndexOf(cValue);
      }
      return nOut;
   }
}

class Program {
   static void Main() {
      var o = new Radix();
      var s = o.encode(1577858399, 36);
      var n = o.decode("q3ezbz", 36);
      Console.WriteLine(s == "q3ezbz" && n == 1577858399);
   }
}
