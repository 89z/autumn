#lang racket
(require net/url)
(copy-port
   (get-pure-port (string->url "http://speedtest.lax.hivelocity.net"))
   (current-output-port)
)