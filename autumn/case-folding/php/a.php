<?php
$s = 'NORTH';
$s = mb_convert_case($s, MB_CASE_FOLD);
var_dump($s == 'north');
