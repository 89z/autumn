import std.file, std.stdio;

void main() {
   auto fname = "unixdict.txt";
   SysTime fileAccessTime, fileModificationTime;
   getTimes(fname, fileAccessTime, fileModificationTime);
   setTimes(fname, fileAccessTime, fileModificationTime);
}
