<?php
$s1 = 'sun mon tue
wed thu fri
';
$s2 = strtok($s1, "\n");
while ($s2 !== false) {
   var_dump($s2);
   $s2 = strtok("\n");
}
