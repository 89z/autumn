#lang racket
(define a
   (current-command-line-arguments)
)
(define s
   (vector-ref a 0)
)
(println
   (string=? s "Sunday")
)
