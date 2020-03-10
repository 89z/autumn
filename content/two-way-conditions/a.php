<?php
$n1 = 10;
# example 1
if ($n1 > 40) {
   var_dump(40);
} else if ($n1 > 30) {
   var_dump(30);
} else {
   var_dump(20);
}
# example 2
if ($n1 > 40) {
   var_dump(40);
} elseif ($n1 > 30) {
   var_dump(30);
} else {
   var_dump(20);
}
# example 3
$n2 = $n1 != 0 ? $n1 : 20;
