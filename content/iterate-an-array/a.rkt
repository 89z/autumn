#lang racket/base
(define a
   (list "Sun" "Mon")
)
(for ((s a))
   (println s)
)
