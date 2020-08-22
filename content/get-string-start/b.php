<?php
extension_loaded('mbstring') or die('mbstring');
$s = mb_substr('αβ', 0, 1);
var_dump($s == 'α');
