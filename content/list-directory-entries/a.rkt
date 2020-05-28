#lang racket/base
(define o1
   (directory-list ".")
)
(for ([o2 o1])
   (println
      (path->string o2)
   )
)
