void main()
{
    import std.algorithm.searching;
    import std.stdio: write, writeln, writef, writefln;
    writeln(countUntil("hello world", "world")); // 6
    writeln(countUntil("hello world", 'r')); // 8
    writeln(countUntil("hello world", "programming")); // -1
    writeln(countUntil("日本語", "本語")); // 1
    writeln(countUntil("日本語", '語')); // 2
    writeln(countUntil("日本語", "五")); // -1
    writeln(countUntil("日本語", '五')); // -1
    writeln(countUntil([0, 7, 12, 22, 9], [12, 22])); // 2
    writeln(countUntil([0, 7, 12, 22, 9], 9)); // 4
    writeln(countUntil!"a > b"([0, 7, 12, 22, 9], 20)); // 3
    
    
}
