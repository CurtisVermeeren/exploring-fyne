# Exploring Fyne
Testing out the Go Fyne toolkit for building cross platform Go applications. https://developer.fyne.io/started/

## Installing Fyne
GCC and graphics library headers are needed to use Fyne with Go. `apt-get install gcc libgl1-mesa-dev xorg-dev` is used in the Dockerfile to install them. 

`go get fyne.io/fyne/v2` Will get the Fyne module. A demo is available using `go run fyne.io/fyne/v2/cmd/fyne_demo` or install it using `go install fyne.io/fyne/v2/cmd/fyne_demo@latest`

## xLaunch
This project uses XLaunch for graphics passthrough. Start a new instance, unselect `Native opengl`, and select `Disable access control`.
The dockerfile and docker-compose contains setup for graphics passthrough

## Packaging with Fyne and Cross-Compiling
Install the fyne cmd using `go install fyne.io/fyne/v2/cmd/fyne@latest`

Use `fyne package -os windows -icon myapp.png` To package an app for windows with an icon

You can also cross compile using Go. `CGO_ENABLED=1 GOOS=windows CC="x86_64-w64-mingw32-gcc" go build .` First install `apt-get install gcc-mingw-w64-x86-64 -y`
