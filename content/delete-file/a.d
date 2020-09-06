void main()
{
    import std.file;
    import std.stdio: writeln, writef, writefln;
    import std.exception : assertThrown;
    
    deleteme.write("Hello");
    writeln(deleteme.readText); // "Hello"
    
    deleteme.remove;
    assertThrown!FileException(deleteme.readText);
    
    
}
