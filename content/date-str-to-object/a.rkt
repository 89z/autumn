#lang racket/base
(define s "2019-12-31")
(require srfi/19)
(string->date s "~Y-~m-~d")
