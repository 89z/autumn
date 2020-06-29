#lang racket/base
(define a '(11 12))
; example 1
(define a1
   (append a '(13 14))
)
(println a1)
; example 2
(define a2
   (cons 10 a)
)
(println a2)
