<?php
$a = ['/', '📗'];

# example 1
$s = json_encode($a);
echo $s, "\n";

# example 2
$s = json_encode($a, JSON_PRETTY_PRINT);
echo $s, "\n";

# example 3
$s = json_encode($a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $s, "\n";
