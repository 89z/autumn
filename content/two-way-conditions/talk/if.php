<?php
$n1 = 10;
# boolean
if ($n1 == 0) { # false
   $n2 = 20;
} else {
   $n2 = $n1;
}
# number
if ($n1) { # true
   $n3 = $n1;
} else {
   $n3 = 20;
}
# print
var_dump($n2, $n3); # 10 10
