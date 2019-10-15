# rBST †
## red/black Binary Search Tree

I admit the name is redundant because all red/black trees are already binary search trees. It is difficult to resist the temptation of a cute name.

I believe the API offered here is roBuST, but the implementation might not be. I did my best to copy Wikipedia, but my grade-school teachers told me Wikipedia can't be trusted, so this program might burn your computer. Please feel free to file issues against this project with abandon.

† [No significant difference](https://en.wikipedia.org/wiki/Bovine_somatotropin#Regulation) has been shown between the [super cow powers](https://unix.stackexchange.com/a/92186) of programs importing rBST-like and non-rBST like packages.

## Motivation

I made this project just to prove a point.

It is often claimed that Go1's lack of generics makes code reuse difficult and forces developers to spend large amounts of time tediously rewriting standard data structures to fit their use case. To the extent that this might be true, I believe it's a net positive that developers might spend more time thinking about how and whether a data structure fits their use case before reaching for that data structure. However, many of the details around this claim are often false.

This project is intended to stand as a counterexample to the broader claim that general-purpose data structures in Go1:
- Can't exist at all
- Can't be made type-safe at compile time
- Can't be supported without relying on packages like `reflect`
- Can't be adapted without code generation
- Are tedious to read, write, or maintain
- Are difficult to understand or use effectively

I do not claim that this is the best-designed red/black tree in Go, the most performant, etc. I believe I have sufficiently tested and shown that it _is_ a red/black tree, and that it works (if not necessarily _well_) with arbitrary values.

While I remain personally unconvinced that parametric types (or any other form of generic programming Go1 doesn't already have) would be a net positive for Go2, I also do not claim that this project proves anything definitive one way or another about the broader topic. The intent here is solely to _disprove_ a common misconception that is often used as an argument in favor of generics; that argument may take other forms, some of which may be valid, and of course other arguments exist that do not depend on any of the above assumptions.
