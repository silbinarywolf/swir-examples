# SWIR Examples

[![Build Status](https://travis-ci.com/silbinarywolf/swir-examples.svg?branch=master)](https://travis-ci.com/silbinarywolf/swir-examples)
[![Documentation](https://godoc.org/github.com/silbinarywolf/swir-examples?status.svg)](https://godoc.org/github.com/silbinarywolf/swir-examples)
[![Report Card](https://goreportcard.com/badge/github.com/silbinarywolf/swir-examples)](https://goreportcard.com/report/github.com/silbinarywolf/swir-examples)

This repository contains examples on how to use the [SWIR package](https://github.com/silbinarywolf/swir).
The reason for seperating the repositories is so that the SWIR package has less dependencies and isn't coupled to Ebiten.

# How to use

1) Build and run the example game (a black screen with red square)
```
go build && ./squaregame
```

2) Move around with the arrow keys and then press the 1 key to write the recording file out. This will override the "record.swirf" file in the "squaregame/test" folder.

3) Run automated input recording test by running:
```
go test ./...
```
