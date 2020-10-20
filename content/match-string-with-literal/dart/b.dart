bool startsWith(String s, String c) {
   return s[0] == c;
}

void main() {
   var s = 'June';
   var b = startsWith(s, 'J');
   print(b);
}
