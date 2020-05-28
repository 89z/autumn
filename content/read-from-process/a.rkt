#lang racket
(define f
   (lambda () (system "ag -V"))
)
(define s
   (with-output-to-string f)
)
(display s)
