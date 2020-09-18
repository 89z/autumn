import std.container.rbtree, std.stdio;

void main() {
   // example 1
   auto t1 = redBlackTree(10);
   t1.insert(11);
   // example 2
   auto t2 = redBlackTree!string("May");
   t2.insert("June");
   // print
   writeln(t1, t2);
}
