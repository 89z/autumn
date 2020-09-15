void main() {
   var a = ['May', 'June'];
   var f = (String s_acc, String s_cur) => s_acc + s_cur;
   var s = a.reduce(f);
   print(s);
}
