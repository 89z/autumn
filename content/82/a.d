void main()
{
    import std.algorithm.iteration;
    import std.stdio: write, writeln, writef, writefln;
    import std.algorithm.comparison : max, min;
    import std.math : approxEqual;
    import std.range;
    
    int[] arr = [ 1, 2, 3, 4, 5 ];
    // Sum all elements
    auto sum = reduce!((a,b) => a + b)(0, arr);
    writeln(sum); // 15
    
    // Sum again, using a string predicate with "a" and "b"
    sum = reduce!"a + b"(0, arr);
    writeln(sum); // 15
    
    // Compute the maximum of all elements
    auto largest = reduce!(max)(arr);
    writeln(largest); // 5
    
    // Max again, but with Uniform Function Call Syntax (UFCS)
    largest = arr.reduce!(max);
    writeln(largest); // 5
    
    // Compute the number of odd elements
    auto odds = reduce!((a,b) => a + (b & 1))(0, arr);
    writeln(odds); // 3
    
    // Compute the sum of squares
    auto ssquares = reduce!((a,b) => a + b * b)(0, arr);
    writeln(ssquares); // 55
    
    // Chain multiple ranges into seed
    int[] a = [ 3, 4 ];
    int[] b = [ 100 ];
    auto r = reduce!("a + b")(chain(a, b));
    writeln(r); // 107
    
    // Mixing convertible types is fair game, too
    double[] c = [ 2.5, 3.0 ];
    auto r1 = reduce!("a + b")(chain(a, b, c));
    assert(approxEqual(r1, 112.5));
    
    // To minimize nesting of parentheses, Uniform Function Call Syntax can be used
    auto r2 = chain(a, b, c).reduce!("a + b");
    assert(approxEqual(r2, 112.5));
}
