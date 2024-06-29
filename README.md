# blume
If it is easy to think of, it should be easy to be programmed.

Blume is a collection of loosely related libraries. The primary goal is to provide a higher level of abstraction over various repetitive, mundane, boilerplate, and commonly used functionalities.

Blume does not follow, try to follow, or claim to follow "best practices". Instead, Blume follows a consistent internal logic and structure.

Blume focuses on low-indirection abstractions. In other words, most functions are trivially inlineable, as there is only a single level of indirection. The point is to make things simpler, not magical. Write procedural code with declarative patterns.

Blume, in effect, creates a unified ecosystem of combinator libraries and eDSLs. That said, every library in Blume is completely functional as a standalone library or tool. Blume exists to maximize reuse and minimize lockin.


## Layers
Blume consists of increasing levels of abstraction and composition.
We will define as such:
- Pure : Libraries which have no (IO) dependencies, hence they are "pure".
- Wrap : Libraries which provide wrappers for existing functionality.
- Comp : Libraries which compose Pure, Wrap, or other Comp libraries.
    - Many comp libraries are also accessible via CLI binaries.

- Pure
    - gen : generics, Map, Reduce, etc.
    - arr : arrays, Any, First, Has, etc. 
    - num : numbers, Abs, Clamp, SameSign, etc.
    - str : strings
    - sar : string arrays
    - val : string value parsing, numbers, dates, time, colors, etc. includes humanized formatting
- Wrap
    - fsio  : filesystem wrapper, provides countless declarative abstractions, combinator library
    - media : comprehensive media related functions, operations, patterns, etc.
    - clog  : common logger, provides an expressive and configurable structured logging framework
    - proc  : cross platform process related functionality
    - info  : cross platform system information
    - comms : simple cross platform IPC library (gRPC based named pipes)
    - srvc  : services and daemons
- Comp
    - parse : a configurable grammarless parser
    - slice : VHLL for generic array operations
    - moe   : a string array combinator DSL
    - list  : file system query DSL (best scaling fs tool, 2x-100,000x faster for 10m+ file dirs)
    - shell : POSIX shell. Interactive or runtime (runtime pipes faster than bash, zsh, pwsh, etc.)
    - syst  : integrated system. jobs, daemons, tasks, IO buses (responsible for comms), etc.
    - eval  : simple expression evaluator, math, input parsing, boolean logic, etc.
    - gs    : Go script, a hodgepodge of functionality
    - blush : Blume shell. Integrates all parts of Blume that make sense to integrate
