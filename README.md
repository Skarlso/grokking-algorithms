# grokking-algorithms

Code for the book [Grokking Algorithms](https://www.amazon.com/Grokking-Algorithms-illustrated-programmers-curious/dp/1617292230) written in Go using Type Parameters ( Generics ).

## What are the additions to the original repository

I've been using type parameters to implement most of the algorithms. From time to time they offer
little purpose other than seeing how they work and how to use them in a real implementation.

The most involved example is [Dijkstra](chapter07/dijkstra.go), which uses the graph outlined in 
type parameters document section [Mutually referencing types](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#mutually-referencing-type-parameters). I do not recommend using this for anything really. But it
was hell of lot fun learning to make this actually work.

I also added a bunch of tests for everything and in some cases, some alternative implementations too.

I will also provide some Benchmarking tests for the heck of it.

I added a bunch of comments to all the code for those who just would like to read the code,
and understand why things are where they are. If anything is missing, or more explanation is required,
please don't hesitate to open an issue.

## Why?

For the glory of the Sontaran Empire of course! That, and because learning is always fun. All in all
the book was a fun little read, and the explanations are great. The author walks you through all the
algorithms step by step to make sure you understand every aspect of it.
