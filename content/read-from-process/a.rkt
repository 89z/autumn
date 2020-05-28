#lang racket ; system
(define f
   (lambda () (system "ag -V"))
)
(define s
   (with-output-to-string f)
)
(display s)
