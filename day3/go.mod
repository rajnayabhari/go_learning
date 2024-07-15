module helloworld

go 1.22.2

require github.com/myuser/calculator v0.0.0

replace github.com/myuser/calculator => ../calculator
// The replace keyword specifies to use a local directory instead of a remote location for the module.
//  In this case, because the helloworld and calculator programs are in $GOPATH/src,
//   the location is simply ../calculator.
//  If the module's source is in a different location, you define the local path here.
