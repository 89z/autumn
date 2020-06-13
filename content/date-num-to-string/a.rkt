#lang racket/base
(require racket/date)
(define n
   (* 366 24 60 60)
)
(define o
   (seconds->date n)
)
(date->string o)
