// example 1
bool f1(String s, String c) {
   return s[0] == c;
}
// example 2
bool f2(String s, String c) => s[0] == c;
// print
void main() {
   var b1 = f1('June', 'J');
   var b2 = f2('June', 'J');
   print(b1 && b2);
}
