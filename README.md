# git resolve-submodule

A stupid simple git plugin written in Go that helps navigate around submodules
like filesystem directories.

## Usage

`$ git resolve-submodule <pathish>`

The pathish value can be ".", ".." or a string that identifies the submodule.

## Installation

1. Install Go and make sure that go bin path (environmental setting) is in $PATH
1. Install git-resolve-submodule (at this point you should be able to run `git resolve-submodule something`
1. Add a function to your dot pile that changes directory to the returned value:

```
function cm {
  cd `git resolve-submodule $1`
}
```

## How to use it

### Change directory to a submodule root

```
$ cm build
```

### Change directory to the root of the parent repository

```
$ cm ..
```

### Changes directory to the root of the current repository

```
$ cm .
```

## How do I identify a submodule with _pathish_?

Use any string that appears in `git submodule`, including the commit id, the path, or the tag. The first
submodule in which the provided "pathish" is found is the submodule which it will cd to.

## Works on:

OSX 10.9.3 / Go 1.3 / Git 1.9.1

I haven't tried anything else.
