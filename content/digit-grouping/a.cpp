#include <iostream>

int main() {
   auto n = 7654321;
   // example A
   std::string sA = std::to_string(n);
   int nA = sA.length() - 3;
   while (nA > 0) {
      sA.insert(nA, ",");
      nA -= 3;
   }
   // example B
   std::string sBa = std::to_string(n);
   std::string sBb;
   int nB = sBa.length();
   for (auto sBc: sBa) {
      if (nB % 3 == 0) {
         sBb += ",";
      }
      sBb += sBc;
      nB--;
   }
   // print
   std::cout << (sA == "7,654,321" && sBb == "7,654,321") << std::endl;
}
