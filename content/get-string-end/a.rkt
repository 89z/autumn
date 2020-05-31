#lang racket/base
(define s "Sunday")
(define s1
   (substring s 1)
)
(println
   (string=? s1 "unday")
)
