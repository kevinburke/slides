talk goal: how can you get started contributing to go?

- background: have been writing Go for three or four years. started contributing
  to the Go project in the last year.

- What does it mean to contribute? What is a contribution? There's this notion
that it's just "the standard library" or "the compiler". Which is wrong. The Go
project is really large and people contribute in a lot of different ways. Here
are some examples of ways to contribute:

    - Add documentation or examples for functions. This is really crucial to
      helping people understand how to use Go and there are a lot of places for
      this.

    - Respond to bug reports. A lot of bug reports are confusing, and it can
      require a few rounds of back and forth to get right. It is really, really
      helpful to ask clarifying questions and try to isolate the code in
      question.

    - Write up issues when you run into problems with Go. A bug report has three
      basic parts:

        - What were you trying to do?
        - What did you expect to see?
        - What did you see instead?

        Don't just limit this to code. Documentation problems are errors too; "I
        couldn't figure out how to use this thing" should be treated as a bug,
        if you gave it your best effort, searched Google etc. and couldn't find
        anything.

        Also: process problems! "I was trying to figure out X and couldn't" is
        helpful, for example Gerrit has been a big problem.

        Try to use Go! The more you use Go the more you will find things that
        are bad about it. That you can report, or fix.

    - Write about your experience using Go! You can make things better for
      everyone by sharing something that you learned.

    - Improve the surrounding infrastructure. There are a lot of parts of the Go
      infrastructure that could function better, or be more pleasant to use.
      We have a bot called gopherbot. It could do a lot more. That could be
      a useful area to contribute!

- How do you contribute? This is sort of tricky. There's no like license for
  contributing, you just get started and doing it. But I can give a few
  suggestions.

Lower your expectations for your first contributions. About 13 out of my first
15 commits to the Go project were for typo fixes and for things like changing
http links to HTTPS. That is totally normal and expected and fine. And what
those commits do is help build up trust that you are going to make positive
changes.

Here are some ideas for finding a way to contribute:

    - Watch the Github issues list. This is the list of new things coming in. You
    might see someone might post an issue that seems like it has an easy fix, in
    which case you can give a comment! Something like "I would like to fix this
    but I am new to contributing and it might take me a while/I might need to ask
    questions about it."

    What's a good candidate for an area to contribute? It might be useful to
    think about what programming area you are either most skilled at, or most
    skilled relative to the average Go developer. For example, I am not very
    good at web development, but I am pretty good relative to most of the
    Go contributors. So this is an area where I try to contribute.

    Your area might be different - maybe security. One area that always needs
    help is the rarer operating systems - OpenBSD, Mac, Solaris, Dragonfly.

- Subscribe to the golang-dev mailing list. (Explain the Go development
cycle). At the start of a new dev cycle there's a thread for 'What do you want
to work on?' There can be useful ideas there, or in other thoughts people post
to the mailing list.

- Try to find an issue and then fix it! Try to notice more when something -
anything - doesn't work the way that you expect it to, and then dig into it. Why
doesn't it work the way you want it to? If you notice it's broken, it's likely
it's broken for other people too!

Math interlude: If there are 750,000 Go developers, and 2% of them use the
package you find a problem with, and 1% of package users run into the problem
you ran into, that's still 150 people! If you fix a problem you can fix it for
150 people which would be pretty great.

    This was one of my first commits to Go. When you copy text using the godoc
    source code viewer, it used to copy the line numbers too. Which would be
    really annoying because you'd have to delete them from wherever you pasted
    them to. So I figured I could fix that, and I did!

- Read through the recent commits. These can give you a good introduction to the
  parts of the Go tools, the types of commit messages people write, and also
  what seem like interesting areas for contribution.

Code contributions. If you want to make a code contribution or docs or examples,
you'll have to learn how to use Gerrit. Gerrit is really hard. Especially if you
are not familiar with Git, or are used to Github where the workflow is usually
just add more commits and push and merge.

We put together a practice repository: if you go to https://bit.ly/goscratch,
there's a scratch repository where you can commit anything, and we'll review it
like it's a real commit. So you can submit whatever you want and screw up as
many times as you want and it's totally fine.

The first contribution is the toughest because you have to install all of the
tools. Once you have all of the tools, it gets easier to do more contributions.

- talk about throwing away work

truman show clip: https://www.youtube.com/watch?v=_Y0jRiIwMUQ
57:46
