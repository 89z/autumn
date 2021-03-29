<?php
$a = ['west', 'east'];
# example 1
foreach ($a as $s) {
   echo $s, "\n";
}
# example 2
foreach ($a as $n => $s) {
   echo $n, "\t", $s, "\n";
}
