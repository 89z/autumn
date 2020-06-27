#lang racket/base
(define n 10)
(if (> n 12)
   "Tue"
   (if (> n 11)
      "Mon"
      "Sun"
   )
)
