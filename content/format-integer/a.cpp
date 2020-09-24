std::string s2 = std::to_string(7654321);
int n2 = s2.length() - 3;
while (n2 > 0) {
   s2.insert(n2, ",");
   n2 -= 3;
}
