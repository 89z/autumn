<?php
declare(strict_types = 1);

$s1 = '!^/prefix/!';
$s2 = $_SERVER['REQUEST_URI'];
$n1 = preg_match($s1, $s2);

if ($n1 == 1) {
   $s3 = preg_replace($s1, '/', $s2);
   header('Location: ' . $s3);
} else {
   return false;
}
