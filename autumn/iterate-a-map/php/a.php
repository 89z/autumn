<?php
$m = ['month' => 12, 'day' => 31];

echo "example 1\n";
foreach ($m as $n) {
   echo $n, "\n";
}

echo "example 2\n";
foreach ($m as $s => $n) {
   echo $s, ':', $n, "\n";
}
