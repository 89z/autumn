<?php
extension_loaded('mbstring') or die('mbstring');
$s = mb_substr('📕📙📒📗', -2);
var_dump($s == '📒📗');
