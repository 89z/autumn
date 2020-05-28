#lang racket ; string-prefix
(define s "Sunday")
; example 1
(string-prefix? s "Su")
; example 2
(string-contains? s "un")
; example 3
(string-suffix? s "ay")
