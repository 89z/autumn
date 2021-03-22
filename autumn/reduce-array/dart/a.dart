void main() {
   var a = ['May', 'June'];
   var s = a.reduce((s, t) => s + t);
   print(s == 'MayJune');
}
