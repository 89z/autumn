#lang racket
(require json)
(define m
   #hash((Sunday . 10))
)
(write-json m)
