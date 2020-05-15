<?php
$s1 = 'https://example.com/sun"mon';
$s2 = escapeshellarg($s1);
var_dump($s2 == '"https://example.com/sun mon"');
