<?php
$m = ['month' => 'May', 'day' => 'Friday'];
$s = http_build_query($m);
var_dump($s == 'month=May&day=Friday');
