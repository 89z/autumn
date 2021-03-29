<?php
$a = ['west', 'east'];

# example 1
for ($n = 0; $n < count($a); $n++) {
   echo $n, "\t", $a[$n], "\n";
}

# example 2
for ($n = 0; key_exists($n, $a); $n++) {
   echo $n, "\t", $a[$n], "\n";
}
