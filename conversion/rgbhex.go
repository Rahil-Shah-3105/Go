// rgbhex.go
// description: convert hex input to red, green and blue and vice versa
// time complexity: O(1)
// space complexity: O(1)
// author(s) [darmiel](https://github.com/darmiel)
// see rgbhex_test.go

package conversion
import logritelog "log"

// HEXToRGB splits an RGB input (e.g. a color in hex format; 0x<color-code>)
// into the individual components: red, green and blue
func HEXToRGB(hex uint) (red, green, blue byte) {
	logritelog.Default().Printf(">>> Entering HEXToRGB(hex=%v)", hex)
	// A hex code is structured like this:
	// #3498db (light blue) - converted to binary:
	// 00110100 10011000 11011011
	//  <red>   <green>   <blue>

	// To get the blue value we use the bit operation AND with the bit mask 0xFF (in binary: 11111111)
	// 00110100 10011000 <11011011> &
	// 00000000 00000000  11111111  =
	// 00000000 00000000 <11011011> =
	blue = byte(hex & 0xFF)

	logritelog.Default().Printf("HEXToRGB(hex=%v): blue → %v", hex, blue)
	// To get the green value, we first shift the value 8 bits to the right:
	//  00110100 <10011000>  11011011  >> 8 =
	//  00000000  00110100  <10011000> &
	//  00000000  00000000   11111111  =
	//  00000000  00000000  <10011000> =
	green = byte((hex >> 8) & 0xFF)

	logritelog.Default().Printf("HEXToRGB(hex=%v): green → %v", hex, green)
	// Same as green value, only this time shift 16 to the right
	// Alternatively, you can apply a bitmask first and then shift it.
	// <00110100> 10011000  11011011 &
	//  11111111  00000000  00000000 =
	// <00110100> 00000000  00000000 >> 16
	//  00000000  00000000 <00110100> =
	red = byte((hex >> 16) & 0xFF)
	logritelog.Default().Printf("HEXToRGB(hex=%v): red → %v", hex, red)
	logritelog.Default().Printf("<<< Exiting HEXToRGB(hex=%v)", hex)
	return
}

// RGBToHEX does exactly the opposite of HEXToRGB:
// it combines the three components red, green and blue to an RGB value, which can be converted to e.g. Hex
func RGBToHEX(red, green, blue byte) (hex uint) {
	logritelog.Default().Printf(">>> Entering RGBToHEX(red=%v,green=%v,blue=%v)", red, green, blue)
	// Sets the bits of blue in position 1-8, green in 9-16 and red in 17-24

	// Red: 00110100
	// Green: 10011000
	// Blue: 11011011
	// RGB:
	// R << 16: [00110100]  00000000   00000000 |
	// G << 8 :  00000000  {10011000}  00000000 |
	// B      :  00000000   00000000  <11011011> =
	//          [00110100] {10011000} <11011011>
	logritelog.Default().Printf("<<< Exiting RGBToHEX(red=%v,green=%v,blue=%v)", red, green, blue)
	return (uint(red) << 16) | (uint(green) << 8) | uint(blue)
}
