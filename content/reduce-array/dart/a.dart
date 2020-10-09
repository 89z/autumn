void main() {
   var a = ['May', 'June'];
   var s = a.reduce((s, s1) => s + s1);
   print(s == 'MayJune');
}
