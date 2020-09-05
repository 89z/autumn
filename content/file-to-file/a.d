void main()
{
    import std.file;
    import std.stdio: writeln, writef, writefln;
    auto source = deleteme ~ "source";
    auto target = deleteme ~ "target";
    auto targetNonExistent = deleteme ~ "target2";
    
    scope(exit) source.remove, target.remove, targetNonExistent.remove;
    
    source.write("source");
    target.write("target");
    
    writeln(target.readText); // "target"
    
    source.copy(target);
    writeln(target.readText); // "source"
    
    source.copy(targetNonExistent);
    writeln(targetNonExistent.readText); // "source"
    
    
}
