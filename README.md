# omacode-golang-present

Presentation on the language Go given to OMG-Code group in Omaha, NE on 2013-12-26.


## Slides
You can view the prsentation via [godoc](http://go-talks.appspot.com/github.com/jamesharr/omacode-golang-present/presentation/presentation.slide), or on your own laptop using [go-present](http://godoc.org/code.google.com/p/go.talks/present).
```bash
$ go get code.google.com/p/go.talks/present
$ go get github.com/jamesharr/omacode-golang-present
$ cd $GOPATH/github.com/jamesharr/omacode-golang-present/presentation
$ present
$ open http://127.0.0.1:3999/
```

The slide deck doesn't have a lot along the lines of subtext or narration (I hate reading off of slides). If you're lost, go check out [Rob Pike's 2012 Go Concurrency Patterns presentation](http://www.youtube.com/watch?v=f6kdp27TYZs). I used most of his content/patterns and adapted it to the audience.

## Contrived Example
I had a contrived+simplified example of a LockManager that locks resources identified by name. It's in the 'lockmanager' directory. I'm not particularly happy with this as an example, but it's a pretty simple example that pulled together what we had learned.


## Learning Resources
* [tour.golang.org](http://tour.golang.org) - Learn & Write Go code right in your browser.
* [Go Docs](http://docs.golang.org/doc/) - Pretty straight forward docs in Go.
* [play.golang.org](http://play.golang.org) - Write arbitrary Go code (no network I/O, nothing that runs longer than 30s).
* Web frameworks in Go (there's a bunch, search around)
   * [beego](https://github.com/astaxie/beego) - I like this one's style
   * [revel](http://robfig.github.io/revel/) - Less of a fan, but it's effective. Go

