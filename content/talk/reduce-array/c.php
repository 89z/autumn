<?php
$a = ['May', 'June'];
$s = '';

foreach ($a as $sc) {
   $s .= $sc;
}

var_dump($s == 'MayJune');
