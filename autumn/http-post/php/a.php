<?php
$http['method'] = 'POST';
$http['header'] = 'Content-Type: application/json';

$http['content'] = <<<eof
{
   "videoId": "lTGjFpjOIP8", "context": {
      "client": {"clientName": "WEB", "clientVersion": "1.19700101"}
   }
}
eof;

$con = stream_context_create(['http' => $http]);
$key = '?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8';

echo file_get_contents(
   'https://www.youtube.com/youtubei/v1/player' . $key, context: $con
);
