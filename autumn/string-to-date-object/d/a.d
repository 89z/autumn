void main()
{
    import std.datetime.systime;
    import std.stdio: write, writeln, writef, writefln;
    import std.datetime.date : DateTime;
    import std.datetime.timezone : UTC;
    
    auto st = SysTime.fromISOExtString("2018-01-01T10:30:00Z");
    writeln(st); // SysTime(DateTime(2018, 1, 1, 10, 30, 0), UTC())
    
    
}
