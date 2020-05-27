#lang racket
(define m
   (make-hash
      '(("Sunday" . 10))
   )
)
(hash-set! m "Monday" 11)
(println m)
