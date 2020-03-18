<?php
$s1 = 'Sunday
Monday';
$s2 = strtok($s1, "\n");
while ($s2 !== false) {
   var_dump($s2);
   $s2 = strtok("\n");
}
