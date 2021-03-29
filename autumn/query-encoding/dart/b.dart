void main() {
   var s = 'west=left&east=right';
   var m = Uri.splitQueryString(s);
   print(m);
}
