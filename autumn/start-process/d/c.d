import std.process;

void main() {
   // example 1
   spawnProcess("dust");
   // example 2
   spawnProcess(["dust", "-V"]);
   // example 3
   spawnProcess("dust", null, Config.none, "..");
}
