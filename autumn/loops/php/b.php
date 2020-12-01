<?php
echo "example 1\n";
$n = 10;
while ($n < 20) {
   echo $n, "\n";
   $n++;
}

echo "example 2\n";
$n = 10;
while (true) {
   if ($n > 19) {
      break;
   }
   echo $n, "\n";
   $n++;
}
