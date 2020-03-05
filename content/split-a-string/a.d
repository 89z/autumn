import std.stdio, std.string;
void main() {
   const s1 = "sun mon tue";
   const s2 = `sun mon tue
wed thu fri
`;
   // example 1
   const a1 = s1.split(" ");
   // example 2
   const a2 = s1.split;
   // example 3
   const a3 = s2.splitLines;
   // print
   writeln(a1, "\n", a2, "\n", a3);
}
