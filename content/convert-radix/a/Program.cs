using System;

class Radix {
   string sDigit = "0123456789abcdefghijklmnopqrstuvwxyz";

   public string encode(int nIn, int nBase) {
      int n = nIn / nBase;
      char c = this.sDigit[nIn % nBase];
      return n > 0 ? encode(n, nBase) + c : c.ToString();
   }

   public int decode(string sIn, int nBase) {
      string s = sIn[..^1];
      int n = this.sDigit.IndexOf(sIn[^1]);
      return s != "" ? decode(s, nBase) * nBase + n : n;
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
