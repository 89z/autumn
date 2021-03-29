import std.stdio, urllibparse;

void main() {
   URLSplitResult u;
   u.query = "west=left&east=right";
   u.writeln;
}
