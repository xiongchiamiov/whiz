Whiz
====

An internal link shortener designed for easy installation.

## Purpose

Whiz is a Golang port of [Zoom], and thus shares its motivation: to provide a
url shortener for internal use at a company, for common urls, Slack topics,
links in presentations, and so on.  For instance, you might host it in the same
domain as your workstations and under the hostname 'whiz'; then if a user adds
a shortlink of 'sandwich' to https://cuberule.com/ , another user visiting
`whiz/sandwich` in their browser will be redirected to that Python
documentation page.

Additionally, I thought it would be a useful way for me to learn a bit of Go:

1. It's small but not trivial.
1. I completely understand the problem, so I only need to focus on the new
   language.
1. It does web serving, which is an important evaluation point for my potential
   uses of go.
1. The "build a static binary and drop it in" model of Go fits well with the
   philosophy Zoom was designed around.

[Zoom]: https://github.com/xiongchiamiov/zoom

Features/Non-Features
---------------------

Whiz is designed for a very specific purpose.  It may or may not be what you
want.

Whiz:

* Installs with one command
* Runs with one command
* Trusts all users

Use
---

    [$]> # TBD

Hacking
-------

    [$]> # TBD
