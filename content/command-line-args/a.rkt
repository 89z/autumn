#lang racket ; vector-ref
(define a
   (current-command-line-arguments)
)
(define s
   (vector-ref a 0)
)
(string=? s "Sunday")
