<?php
function f_join(...$a_elem) {
   return implode(DIRECTORY_SEPARATOR, $a_elem);
}
$s_pub = getenv('PUBLIC');
# round 1
$s_old = realpath('abs.php');
$s_new = f_join($s_pub, 'Music', 'abs.php');
symlink($s_old, $s_new);
# round 2
$s_old = realpath('chdir.php');
$s_new = f_join($s_pub, 'Videos', 'chdir.php');
symlink($s_old, $s_new);
