import std.algorithm, std.stdio;

void main() {
   auto a = ["May", "June"];
   // example A
   auto sA = a.reduce!((sY, sZ) => sY ~ sZ);
   // example B
   auto f = (string sY, string sZ) => sY ~ sZ;
   auto sB = a.reduce!(f);
   // print
   writeln(sA == "MayJune" && sB == "MayJune");
}
