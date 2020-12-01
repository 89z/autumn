using C = System.Console;
using J = System.Text.Json.JsonDocument;

class Program {
   static void Main() {
      var in_s = "{\"U2\": {\"Boy\": [\"Twilight\", \"I Will Follow\"]}}";
      var m = J.Parse(in_s);
      var out_s = m.
         RootElement.
         GetProperty("U2").
         GetProperty("Boy")[0].
         ToString();
      C.WriteLine(out_s == "Twilight");
   }
}
