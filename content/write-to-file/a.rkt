#lang racket/base
(define o
   (open-output-file "a.txt")
)
(write-string "Sunday\n" o)
