#lang racket/base
(define s1 "Sunday")
(define s2
   (substring s1 1 3)
)
(string=? s2 "un")
