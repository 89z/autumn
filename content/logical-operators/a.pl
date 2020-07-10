$b_t = 11 > 10;
$b_f = 10 > 11;
# example 1
$b1 = ! $b_f;
# example 2
$b2 = $b_f || $b_t;
# example 3
$b3 = $b_t && $b_t;
# print
print $b1, $b2, $b3, "\n";
