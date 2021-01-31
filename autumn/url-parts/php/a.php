<?php
$m = ['one' => 'odd', 'two' => 'even'];
$s = http_build_query($m);
var_dump($s == 'one=odd&two=even');
