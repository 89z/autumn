void main() {
   var s = 'month=May&day=Friday';
   var m = Uri.splitQueryString(s);
   print(m);
}
