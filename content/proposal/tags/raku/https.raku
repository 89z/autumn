use LWP::Simple;
my $s1 = "https://speedtest.lax.hivelocity.net";
my $s2 = LWP::Simple.get($s1);
print $s2;
