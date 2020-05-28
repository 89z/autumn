#lang racket/base
(define n 0)
(if (positive? n)
   "positive"
   (if (negative? n)
      "negative"
      "zero"
   )
)
