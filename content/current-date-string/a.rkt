#lang racket/base
(require racket/date)
(define o
   (current-date)
)
(date->string o)
