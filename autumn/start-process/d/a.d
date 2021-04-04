import std.process;

void main() {
   // example 1
   spawnProcess("dust");
   // example 2
   ["dust", "-V"].spawnProcess;
}
