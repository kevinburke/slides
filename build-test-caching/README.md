# Build and Test Caching

What is a cache?

Make an extra copy of something to get a result faster.

Faster - generally you cache stuff that is either slow to compute or far away

You use caches everyday:

- CPU caches
- Memcache (database cache)
- Computation (digits of pi)
- CDN
- Browser cache
- Artifacts

Considerations/tradeoffs

- Staleness vs hit rate - data from any copy may be out of sync with the
original. Easier if the original never changes.

- Size vs hit rate - how much space do you have for a cache. How do you evict
items?

Loading data from any copy risks having the data go out of date. This only
matters if the data _can_ go out of date (database, CPU/RAM). In those cases
need to be extremely careful about invalidation

Simple cache algorithm:

1. Compute a cache key
2. Use key to search cache for item
3. If it exists in cache, return it - easy case.
4. Do the slow computation
5. Save the data in the cache at the key
6. Return the data

5a. Make space in the cache for the new item. Different strategies to do this
- evict oldest item, evict least recently used item, most common.

Build caches in particular

Make/Makefile

Make was invented in 1976! It's one of the earliest build tools. The advantage
of Make is it's on every machine and it's very easy to get started with.

I've set up a sample project which is a very simple HTTP server, which loads
a single template. The template is saved as an HTML file and loaded into the Go
program with go-bindata. We're going to use Make to build this project.

Make has a concept of targets, and dependencies that are used to build those
targets.

```
target: dependency1 dependency2 dependency3
	command that builds target from dependencies...
```

Let's call our binary "gophercon." If we wanted to use Make to build gophercon,
we would express "gophercon" as the target, and all of the input files as
dependencies.

```
gophercon: main.go bindata.go
    go build -o gophercon .
```

Run it.

"make gophercon"

Then run it again.

Note when we run it again, Make tells us that the file was up to date, in other
words, it didn't have to do anything.

How does it work? Make looks at the modification times of each file. If the
target is newer than the inputs, Make does not re-run the command. We can check
the modification times with the stat command.

You can also chain dependencies. We can add the go-bindata process as part of
the dependency graph, like this.

```
bindata.go: templates/index.html
	go-bindata -o bindata.go ./templates

gophercon: main.go bindata.go
    go build -o gophercon .

serve: gophercon
    ./gophercon
```

And you can imagine adding more complexity from there.

Downsides of Make:

- If you change the command, you won't get the files to be rebuilt.

- If you check out older version of source, Make won't rebuild target.

- You have to redeclare the dependency graph that Go gives you for free, and you
  might get it wrong.

That's a summary of Make - very easy to get started with.

Bazel is a build system that was written at Google to build large C/C++
applications more quickly, and later open sourced.

Bazel's goals are, number one, reproducibility - everyone should be able to
reproduce the same build artifact, which means being really careful about
removing the differences between environments. Number two, speed. Bazel wants to
avoid doing work if it doesn't have to do it, and it wants to be able to do as
much as possible in parallel.

Here's the Bazel configuration for an example Go project, one of my projects.
It's a completely different way of looking at things from Make.

- Configuration is written in Starlark. An interesting thing about Starlark is
it is guaranteed to terminate, which means it is not Turing complete. This means
you can use if statements but you can't use things like while statements and you
can't use recursion.

- Extremely explicit about dependencies. It's difficult to get started with
Bazel without adding your entire project to Bazel.

We need to declare everything that we're using here and all of the dependencies
- there are tools that help with this but you still need to be explicit.

bazel build -s //...:all

When it runs, Bazel is going to build a dependency graph of my entire project,
and use it to make decisions about what needs to be rebuilt and how much of it
can be done in parallel.

What Bazel is going to do is take everything I've added here and to the max
extent possible, produce a hermetically sealed build.

Point out it is downloading the Go SDK and the JDK. Compiling and running
everything.

Run ls -la

bazel-bin/cmd/nacl-generate-key/nacl-generate-key_/nacl-generate-key -h

How does Bazel cache builds? in bazel-bin, outputs are language dependent. Go
for example uses ".x" files as the cache key.

The .x file is a subset of the data in the .a file, only the parts that change.

Each build bazel will check if the .x file is up to date (including all
dependencies), if not it will recompile the package and generate a new .a file.

Bazel is like an F1 car. It's really powerful but it has a lot of startup cost
and it takes a lot of tuning to get right.

Some of the other disadvantages of Bazel:

A _lot_ of stuff is in the build cache! Including an entire copy of the JDK.
This can make it difficult to save time on systems like Travis CI. To implement
the build cache speedups you really need to put the cache behind a persistent
server. This is trivial for Google, tougher for Github Actions or Circle or
Travis CI.

Now we're going to look at Go, which strikes a balance between these.

Go:

Go build caching uses a lot of the same semantics as Bazel but has the advantage of only needing to build Go programs.

What should we put in the cache? We could cache the entire program, we could cache functions, we could cache files.

Go caches at the _package_ level. This is a natural fit because we already build object files, these .a files, for each package. A Go binary is created by linking all of these .a files together.

Now that we know what we are putting in the cache, we need a cache key. How does Go compute a cache key? Well, any input to a package (or the program) that could generate a different .a file needs to be part of the cache key. If they do change, the cache key must be different, because the .a file will be different, and we need to force Go to rebuild the package.

Basic idea is, concatenate all of the things that could cause the .a file to change and compute a checksum.

What is a checksum? A 32 byte value that makes it really easy to compare whether two documents are the same. Short, change one character and get a different result. Only need to compute the checksum once, can use it for every comparison.

    Simplest possible example: a package with no dependencies. We're going to
    compile this package with a special environment variable that will print out
    what the build cache is doing.

    GODEBUG=gocachehash=1 go install -v .

    So it installed. You can see that the Go build cache is building a hash by
    printing out a bunch of the inputs to the hash. The first line you can see
    the Go version. Below that you can see a bunch of other inputs to the cache
    - the goos, goarch, all of these contain info about how the package was
    built.. At the bottom you have a checksum, which is the checksum of all of
    the previous lines.

    Let's go look for that checksum in the cache. You can find the cache by
    running "go env GOCACHE" on your device. There are 256 folders in the cache
    - it's just the first two letters of the checksum. This way we don't have
    one huge folder with all of the cache files.

    See how the actual file ends in "-a"? That means it's an _action_ file -
    basically a pointer between one key and another key. If we open the file, we
    see our checksum on the left side - that's a checksum of all of the inputs.
    On the right side we see another checksum. That's a checksum of the .a file
    - the output file.

    We can load the output file at its checksum in the cache. It's got a bunch
    of garbage in it because it has a lot of object code in it that's meant to
    be read by the computer.

    - cut -

    Now let's do a version of this package, but it has a dependency on another
    very simple package, the package on the right. How does that work?

    GODEBUG=gocachehash=1 go install -v .

    Go will build the leaf package first, pkg2, in the same way as we built the
    base package before. Let's look at the object file for package 2 in the
    cache - it has some important data in it.

    You can see these two base64 encoded ids in the object file. The left side
    is the checksum of the inputs, re-encoded as base64. The right side is the
    checksum of the object file, sans ID's.

    This base64 sum gets encoded as part of the checksum for the importing
    package, on this line here. This way, if the dependency changes, the base64
    code will change, and the dependent package will get a new checksum, which
    will trigger a rebuild. We can continue on like this until our entire
    program is built - it's turtles all the way down.

    Concluding:

    - Start with the leaf packages
    - Build cache keys for each one
    - Build the packages they depend on.

Go also has the ability to cache test results. How does it do that? Well, test
caching is just like build caching, with a few tweaks.

Let's compile this test program so we can see it in action.

The first thing we can see is that there's a lot more output on the screen. This
is because to run a test, we have to import the testing package, which imports
a lot of packages from the standard library. There's nothing we can do about
that.

If we scroll up, we can find the important part, which is the cache instruction
for our test package. We can see that the package under test is considered an
input to the test package. There's also this file in here called testmain.go,
which is generated by Go and is what's used to actually run your tests.

The next thing we can see is that if we run the test again, the output gets
cached. For now, Go is using the same strategy that it uses in the build, to
decide what gets cached.

However, one thing tests is that they often have side effects. They load files,
or they read environment variables. If you change the timezone in the TZ
environment variable, and re-run the tests, Go should not cache tests that try
to create a time.Time object and do math on it, or you'd be told that your tests
were passing, incorrectly.

- cut -

There are four syscalls that Go monitors for side effects. We can see them all
here, in the guts of the Go source code.

    - open - note only files opened in GOROOT/GOPATH
    - stat
    - chdir
    - getenv - Any time we read an environment variable

    Any time you invoke any of these syscalls, Go writes the names to a special
    log file for the test.

    At the end of the test run, Go reads through the test log and looks up the
    values of every key, for example, we look at the modification time for a
    file we opened. I've modified our test to add two simple syscalls, so you
    can see that in action here.

    (RUN TEST)

    We can look up the testlog by looking at the first test binary hash listed
    here, which contains a pointer to the test log.

    Then you can see we are looking up the values of each of these variables,
    and writing a hash for them. This hash is combined with the testlog, and the
    right side is the successful test output.

    The next time we run the test, how does Go know that the files and
    environment variables match before it runs the test? We do a trick - open
    all of them _before_ the test runs this time, and we check that they have
    the same values as they did the last time we ran the test.

    We can see this when we run the test again. The first time, the test result
    printed before the testInputs and getenv lines. This time, the test result
    is printed *after* those lines.

HASH subkey c08f222deca3e6a283a845c5662b93efe2c5a8ecfb978e31a19043cf63017e46 "inputs:a8f2ef068927abbb05af137326d2eeb9cbdadd9ac57064bf1f50e0192530aae4" = e8b5e24f040aa9baa87160325552c99d1b4b2ae1953dca11a4478a4c51f893b6
HASH subkey (pointer to test log) (hash of test input values) = test result
