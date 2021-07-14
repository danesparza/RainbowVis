# RainbowVis
Port of the [javascript RainbowVis library](https://github.com/anomal/RainbowVis-JS) to Go

# Usage
Install in your project
```shell 
get github.com/danesparza/RainbowVis
```
Then
```go
color1 := "ff000e" // Bright red
color2 := "ff7c00" // Bright orange
color3 := "f6ff00" // Bright yellow

// Create our rainbow type
rb := rainbow.GetRainbow()

// Set the spectrum, and then get the color at the given number in the range
err := rb.SetSpectrum(color1, color2, color3)
v1 := rb.ColorAt(23) // ff3908
```
