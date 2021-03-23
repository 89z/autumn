<?php
$s = 'March';
$s = mb_convert_case($s, MB_CASE_FOLD);
var_dump($s == 'march');
