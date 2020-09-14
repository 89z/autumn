<?php
extension_loaded('mbstring') or die('mbstring');
$s = mb_substr('📕📙📒📗', 0, 2);
var_dump($s == '📕📙');
