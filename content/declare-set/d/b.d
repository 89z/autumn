import std.container.rbtree, std.stdio;

void main() {
   // example 1
   auto t1 = redBlackTree(10, 11);
   // example 2
   auto t2 = redBlackTree!string("May", "June");
   // print
   writeln(t1, t2);
}
