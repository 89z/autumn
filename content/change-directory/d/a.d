void main()
{
    import std.file;
    import std.stdio: writeln, writef, writefln;
    import std.algorithm.comparison : equal;
    import std.path : buildPath;
    
    auto cwd = getcwd;
    auto dir = deleteme ~ "dir";
    dir.mkdir;
    scope(exit) cwd.chdir, dir.rmdirRecurse;
    
    dir.buildPath("a").write(".");
    dir.chdir; // step into dir
    "b".write(".");
    dirEntries(".", SpanMode.shallow).equal(
        [".".buildPath("b"), ".".buildPath("a")]
    );
    
    
}
