void main() {
   var a = ['May', 'June'];
   var n = a.fold(0, (num n, s) => n + s.length);
   print(n == 7);
}
