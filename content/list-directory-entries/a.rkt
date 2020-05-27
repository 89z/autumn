#lang racket
(for ([path (directory-list ".")])
   (println
      (path->string path)
   )
)
