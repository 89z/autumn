<?php
declare(strict_types = 1);


$a = [
   '',
   'date number',
   'date object',
   'date string',
   'duration number',
   'duration string'
];

foreach ($a as $s1) {
   foreach ($a as $s2) {
      foreach ($a as $s3) {
         if ($s3 == '') {
            continue;
         }
         # 1
         if ($s1 == '' && $s2 == '') {
            if (str_starts_with($s3, 'duration')) {
               continue;
            }
            $m['-> ' . $s3] = true;
            continue;
         }
         # 2
         if ($s1 == '' || $s2 == '') {
            if (str_starts_with($s1, 'duration')) {
               continue;
            }
            if (str_starts_with($s2, 'duration')) {
               continue;
            }
            if (str_starts_with($s3, 'duration')) {
               continue;
            }
            if ($s1 == $s3 || $s2 == $s3) {
               continue;
            }
            if ($s1 == '') {
               $m[$s2 . ' -> ' . $s3] = true;
               continue;
            }
            $m[$s1 . ' -> ' . $s3] = true;
            continue;
         }
         # 3
         if ($s1 != 'date object') {
            continue;
         }
         if (str_starts_with($s2, 'duration')) {
            $m['date object, duration -> date object'] = true;
            continue;
         }
         if (! str_starts_with($s3, 'duration')) {
            continue;
         }
         $m['date object, date object -> ' . $s3] = true;
      }
   }
}

foreach ($m as $s => $b) {
   echo $s, "\n";
}
