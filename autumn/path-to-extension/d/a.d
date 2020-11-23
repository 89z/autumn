import std.range : empty;
assert(extension("file").empty);
writeln(extension("file.")); // "."
writeln(extension("file.ext"w)); // ".ext"
writeln(extension("file.ext1.ext2"d)); // ".ext2"
assert(extension(".foo".dup).empty);
writeln(extension(".foo.ext"w.dup)); // ".ext"

static assert(extension("file").empty);
static assert(extension("file.ext") == ".ext");
