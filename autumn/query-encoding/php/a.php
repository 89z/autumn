<?php
$m = ['west' => 'left', 'east' => 'right'];
$s = http_build_query($m);
var_dump($s == 'west=left&east=right');
