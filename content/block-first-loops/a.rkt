#lang racket/base
(define (f n)
   (println n)
   (when (< n 19)
      (f (add1 n))
   )
)
(f 10)
