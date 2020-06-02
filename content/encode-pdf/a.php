<?php
require 'dompdf/autoload.inc.php';
use Dompdf\Dompdf;
# read
$dompdf = new Dompdf();
$dompdf->getOptions()->setIsFontSubsettingEnabled(true);
$s_in = file_get_contents('index.html');
$dompdf->loadHtml($s_in);
# write
$dompdf->render();
$s_out = $dompdf->output();
file_put_contents('index.pdf', $s_out);
