# Bluri
Bluri can apply a blur effect on a image you pass as an argument.
I use bluri to blur my i3-lock screen since ´convert´ was not fast enough.

## Dependencies
To use bluri you will need a working go installation. Also you need to install the package this program is depending on:
```
go get github.com/disintegration/imaging
```

## Build
```
cd bluri
go install .
bluri -f=path/to/image.png -b=blurfactor
```

## Example
![alt text](screenshot.png?raw=true "screenshot.png")