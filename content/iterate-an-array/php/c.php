<?php
$a = ['May', 'June'];
$n = count($a) - 1;

while ($n >= 0) {
   echo $n, "\t", $a[$n], "\n";
   $n--;
}
