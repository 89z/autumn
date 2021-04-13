# Steven Penny

Here are some items I have worked on in April and March 2021. I will arrange the
items by programming language.

## Go

I will start with my contributions to other projects. In March I had three pull
requests merged in `google/vim-ft-go` [1][2][3]. This is the official repository
for Go syntax highlighting for the Vim text editor. I also had one merged pull
request at `alecthomas/chroma` [4]. Chroma provides syntax highlighting for code
blocks in web pages; the code blocks can contain code in many programming
languages.

So far in April, I have had four pull requests merged. I had two merged at
`noborus/ov` [5][6], and two merged at `walles/moar` [7][8]. These projects both
provide a pager for command line programs. So instead of relying on a system
command like `less` or similar, these modules can be used. They allow for paging
lengthy output in a controlled fashion, you can use the keyboard to scroll
through the data at your own speed, without it all flooding out at once. They
also provide support for searching through the text using regular expressions.
I helped to add support for maintaining the output on screen after closing the
program. Finally, this month I also made my first pull request to `golang/go` [9].
This is the repository for the Go programming language. My pull request is for
adding a function to the `path/filepath` package. It would add a function called
`JoinList`, which someone had first requested in an issue in 2018. My patch has
to go through a review process, so I am not sure what will happen yet.

Regarding my own projects, I will start with newly created projects. Last month
I started `89z/deezer` [10]. This is a module and command line tool for working
with the site `deezer.com`. It allows you to interact with at least four Deezer
APIs, with regard to getting metadata for music, to even downloading and
decrypting media streams.

This month, I created a repository called `89z/page` [11]. Inspired by my other
recent work, this is a pager for command line programs. I had issue with some
of the existing projects in this space, as in one case a project had not
created a separate module for their command line tool, so it blew their project
dependencies higher than it needed to be. For my own needs, I want just a small
basic pager with limited features. I prefer to pull in as few third party
dependencies as I can.

Now I will talk a little about my older projects. I should clarify, while they
were not created recently, I am still active on all of these, with commits on
all of them in the last week. First I have `89z/autumn` [12]. This is a website
for comparing programming languages. My goal is to support 10 different
programming languages. Currently I have good support for nine: D, Dart, Go,
JavaScript, Nim, PHP, Python, Ruby and Rust. The site holds many pages of
programming tasks in these languages. For example, on my site the Go language
currently has 78 tasks, each on its own page. On each page, you will usually find
multiple code examples, as well as links to official documentation, and links to
similar help sites. This is very useful to me in learning programming. When I
have a basic programming task, I try to implement it in all the languages. I find
by doing this, I can see weakness in some languages. Also it is useful, in that
trying to create implementations in different languages exposes the language
differences, and results in better code all around. I can then take what I learn
and put it to use with other projects.

Next I have a module simply called `89z/x` [13]. This is a place where I can put
functions that I use with multiple projects. For example, I have a function for
downloading files, that can check the destination already exists. Also it
contains logging functions, that will output program progress in color. Also
included are some helper functions for extracting archives such as `bz2`, `gz`,
`xz`, `zip` and `zst`. Finally I created a type for running external programs.
Go already includes a good `os/exec` package, but it can be verbose and I didnt
care for the API.

Next I have a project called `89z/winter` [14]. This is a music project, that
utilizes SQLite as the backend. I like SQLite, but I was not able to find a pure
Go SQLite module, so eventually I would like to move to a Go database such as
`genjidb/genji` [15] or similar. My project Winter is designed to allow a user
to rate and find new music. It gives a command line output of albums and songs by
an artist, with links to playlists for listening. I have also though about
writing a server so that you could interact with the data using the browser,
instead of command line.

1. https://github.com/google/vim-ft-go/pull/7
2. https://github.com/google/vim-ft-go/pull/8
3. https://github.com/google/vim-ft-go/pull/9
4. https://github.com/alecthomas/chroma/pull/457
5. https://github.com/noborus/ov/pull/67
6. https://github.com/noborus/ov/pull/69
7. https://github.com/walles/moar/pull/39
8. https://github.com/walles/moar/pull/40
9. https://github.com/golang/go/pull/45450
10. https://github.com/89z/deezer
11. https://github.com/89z/page
12. https://github.com/89z/autumn
13. https://github.com/89z/x
14. https://github.com/89z/winter
15. https://github.com/genjidb/genji

## JavaScript

Regarding JavaScript, I just one one project `89z/umber` [1]. This is a website
I made for posting and listening to music. Previously I was using social media
sites like Reddit to post links to music I like. However I found that Reddit
search will only return the 1,000 most recent results, even with special search
operators. So I dumped my data into a JSON file, then the site parses the JSON
file with JavaScript.

Recently I also added support for LocalStorage [2], which keeps track of the
songs that have been listened to, and marks them with a purple color. This is
similar to the browser history, but I wanted something divorced from that, I
want wanted to be able to clear the data, without clearing the browser history.
Also this allows me to expand in the future, for example instead of just boolean
values, I could track counts and implement "frecency" sorting or similar.

1. https://github.com/89z/umber
2. https://developer.mozilla.org/docs/Web/API/Web_Storage_API
