<?php
$o = new DateTime('2019-12-31');
$o2 = new DateTime('2019-12-31T23:59:59');
$s = $o2->diff($o)->format('%H:%I:%S');
var_dump($s == '23:59:59');
