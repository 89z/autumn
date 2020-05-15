<?php
$s1 = 'https://example.com/sun"mon';
$s2 = escapeshellcmd($s1);
system('firefox ' . $s2);
