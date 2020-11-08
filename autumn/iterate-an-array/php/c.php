<?php
$a = ['May', 'June'];
# example 1
$n = 0;
while ($n < count($a)) {
   echo $n, "\t", $a[$n], "\n";
   $n++;
}
# example 2
$n = 0;
while (key_exists($n, $a)) {
   echo $n, "\t", $a[$n], "\n";
   $n++;
}
