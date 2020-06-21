<?php
$s1 = '!^/prefix/!';
$s2 = $_SERVER['REQUEST_URI'];
$n = preg_match($s1, $s2);
if ($n !== 0) {
   $s3 = preg_replace($s1, '/', $s2);
   header('Location: ' . $s3);
} else {
   return false;
}
