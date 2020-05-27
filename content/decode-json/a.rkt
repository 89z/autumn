#lang racket
(require json)
(define s #<<eof
{"Sunday": 10}
eof
)
(string->jsexpr s)
