#lang racket
(define m (make-hash))
(hash-set! m "Sunday" 10)
(hash-ref m "Sunday")
