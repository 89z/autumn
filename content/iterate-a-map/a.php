<?php
$m1 = ['Sun' => 10, 'Mon' => 11];
# example 1
foreach ($m1 as $n1) {
   var_dump($n1);
}
# example 2
end($m1);
while (true) {
   $s1 = key($m1);
   if ($s1 === null) {
      break;
   }
   $n1 = current($m1);
   var_dump($s1, $n1);
   prev($m1);
}
