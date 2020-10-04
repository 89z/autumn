using Newtonsoft.Json.Linq;
using System.IO;
using System;

class Program {
   static void Main() {
      JObject x = JObject.Parse(File.ReadAllText("data.json"));
   }
}
