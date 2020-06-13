#lang racket/base
(define s_f_red "\e[31m")
(define s_end "\e[m")
(displayln
   (string-append s_f_red "Sunday" s_end)
)
