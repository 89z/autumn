import std.container.rbtree, std.stdio;

void main() {
   auto t = redBlackTree(10);
   t.insert(11);
   t.writeln;
}
