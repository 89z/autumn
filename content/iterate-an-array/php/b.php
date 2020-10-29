<?php
$a = ['May', 'June'];
$f = fn ($s) => var_dump($s);
array_walk($a, $f);
