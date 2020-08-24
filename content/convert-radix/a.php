<?php

class Radix64 {
   private $s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';

   function Encode(int $n_in): string {
      $s_out = '';
      do {
         $s_out = $this->s_dig[$n_in % 64] . $s_out;
         $n_in = intdiv($n_in, 64);
      } while ($n_in > 0);
      return $s_out;
   }

   function Decode(string $s_in): int {
      $n_out = 0;
      $a_in = str_split($s_in);
      foreach ($a_in as $s_chr) {
         $n_out = $n_out * 64 + strpos($this->s_dig, $s_chr);
      }
      return $n_out;
   }
}

$n = time();
$o = new Radix64;
$s = $o->Encode($n);
$n2 = $o->Decode($s);
var_dump($n, $s, $n2 == $n);
