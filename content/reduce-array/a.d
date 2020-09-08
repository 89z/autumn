import std.stdio, std.algorithm, std.typecons;
void main() {
    immutable array = [1, 2, 3, 4, 5];
    immutable r = reduce!(q{a + b}, q{a * b})(tuple(0, 1), array);
    writeln("Sum: ", r[0]);
    writeln("Product: ", r[1]);
}
