#lang racket
(define a1 '(10 11))
(define a2
   (append a1 '(12 13))
)
(writeln a2)
