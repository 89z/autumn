<?php
$a = ['May', 'June'];

echo "example 1\n";
foreach ($a as $s) {
   echo $s, "\n";
}

echo "example 2\n";
foreach ($a as $n => $s) {
   echo $n, "\t", $s, "\n";
}
