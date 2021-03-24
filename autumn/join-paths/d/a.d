void main()
{
    import std.path;
    import std.stdio: write, writeln, writef, writefln;
    version (Posix)
    {
        writeln(buildPath("foo", "bar", "baz")); // "foo/bar/baz"
        writeln(buildPath("/foo/", "bar/baz")); // "/foo/bar/baz"
        writeln(buildPath("/foo", "/bar")); // "/bar"
    }
    
    version (Windows)
    {
        writeln(buildPath("foo", "bar", "baz")); // `foo\bar\baz`
        writeln(buildPath(`c:\foo`, `bar\baz`)); // `c:\foo\bar\baz`
        writeln(buildPath("foo", `d:\bar`)); // `d:\bar`
        writeln(buildPath("foo", `\bar`)); // `\bar`
        writeln(buildPath(`c:\foo`, `\bar`)); // `c:\bar`
    }
    
    
}
