void main() {
   var a = ['May', 'June'];
   // example A
   var sA = a.reduce((sY, sZ) => sY + sZ);
   // example B
   var f = (String sY, String sZ) => sY + sZ;
   var sB = a.reduce(f);
   // print
   print(sA == 'MayJune' && sB == 'MayJune');
}
