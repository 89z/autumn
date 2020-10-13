import std.getopt;

void main(string[] args)
{
    void bat(){} void fox(){}
    getopt(args, "b", &bat, "f", &fox);
}
