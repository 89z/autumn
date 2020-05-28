#lang racket ; string-join
(define a
   (list "Sun" "Mon")
)
; example 1
(define s1
   (string-join a ",")
)
(string=? s1 "Sun,Mon")
; example 2
(define s2
   (string-join a)
)
(string=? s2 "Sun Mon")
