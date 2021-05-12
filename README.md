## Cynthia
This is a tool for TDD(Test Driven Development).

Cynthia ensures test-driven development. When you use it in CI, it ensures that untested methods are not introduced.
The reason for this is that Cynthia alerts you to a function that has no tests.

### derivation
The Greek pronunciation of the word `συνήθεια`, meaning habit, converted to the alphabet.

The most important thing in test-driven development is to make it a habit.
By writing from the test first, you'll have a clean, easy-to-use function.
We hope that you'll get into the habit of doing so.
That's why I named it after this tool.

## Getting Stated
### How to Install

```
go get github.com/riita10069/cynthia/cmd/cynthia
```

### Usage Example
```
go vet -vettool=`which cynthia` [package path]
go vet -vettool=`which cynthia` [package path1] [package path2] [package path3]
```
