void main()
{
    import std.conv;
    import std.stdio: write, writeln, writef, writefln;
    import std.algorithm.comparison : equal;
    
    assert(toChars(1).equal("1"));
    assert(toChars(1_000_000).equal("1000000"));
    
    assert(toChars!(2)(2U).equal("10"));
    assert(toChars!(16)(255U).equal("ff"));
    assert(toChars!(16, char, LetterCase.upper)(255U).equal("FF"));
    
    
}
