toHumanReadable(int size) {
  final units = "KMGT";
  var unit = -1;
  var left = size * 1.0;
  while (left > 1100 && unit < 3) {
    left /= 1024;
    unit++;
  }
  if (unit == -1) {
    return "${size}B";
  } else {
    return "${left.toStringAsFixed(1)}${units[unit]}iB";
  }
}

void main() {
   var s = toHumanReadable(1264);
   print(s);
}

