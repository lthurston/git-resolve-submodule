# git resolve-submodule

A stupid simple git plugin written in Go that helps navigate around submodules
like filesystem directories.

## Usage

`$ git resolve-submodule <pathish>|-v`

The pathish value can be ".", ".." or a string that identifies the submodule.

The `-v` or `--version` flags do what you think.

## Installation

1) Install Go (`brew install go`, or whatever)

2) Set up a go dev environment: `mkdir -p ~/go/{src,bin}`

3) Add some stuff to your .bash_profile:
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

4) Get thee to the Go-ery: `cd ~/go`

5) Fetch the source: `go get github.com/lthurston/git-resolve-submodule`

6) Compile and install: `go install src/github.com/lthurston/git-resolve-submodule/git-resolve-submodule.go `

7) Add a function to your dot pile that changes directory to the returned value:

```
function cm {
  path=`git resolve-submodule $1`
  if [ $? -eq 0 ]; then
    cd $path
  else
    echo 'No se puede.'
  fi
}
```


## How to use it

### Get thee to the top level of a submodule

```
$ cm stuff
```

Where `stuff` is a string that appears in the output of `git submodule`.

### Get thee to the top level of the containing repository

```
$ cm ..
```

### Get thee to the top level of the current repository / submodule

```
$ cm .
```

### Get thee to the outermost git repository

```
$ cm /
```

## How do I identify a submodule with _pathish_?

Use any string that appears in `git submodule`, including the commit id, the path, or the tag. The first
submodule in which the provided "pathish" is found is the submodule which it will cd to.

## Can I get cray-see with my _pathish_, like `cm build/../..`?

Not right now, but it's possible. Dunno if it'd be all that useful though. I do see the advantage of `cm /` which would take you to the topmost repo if you're nusted and nuzzled in there real deep. Maybe I'll implement that.


## Works on:

OSX 10.9.3 / Go 1.4 / Git 1.9 - 2.2

I haven't tried anything else.
