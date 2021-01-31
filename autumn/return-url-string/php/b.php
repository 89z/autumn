<?php
$s = 'one=odd&two=even';
parse_str($s, $m);
var_dump($m);
