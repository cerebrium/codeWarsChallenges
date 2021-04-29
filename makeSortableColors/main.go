package main

import (
	"fmt"
	"strconv"
)

type Color struct {
	color string

	name string

	_id string
}

// structure for the color object thats sortable
type ColorObject struct {
	_id string

	color string

	name string

	blue int16

	red int16

	green int16

	chroma float64

	sat float64

	val float64

	luma float64

	hue float64
}

// array maker helper
func createColor(color Color) *ColorObject {
	// make the object to return
	m := new(ColorObject)

	// take the color and generate red, blue and green
	r, err := strconv.ParseInt(color.color[1:3], 16, 64)
	g, err := strconv.ParseInt(color.color[3:5], 16, 64)
	b, err := strconv.ParseInt(color.color[5:7], 16, 64)

	// check for error
	if err != nil {
		fmt.Println(err)
	}

	// divide and convert type
	r2 := float64(r) / float64(255)
	g2 := float64(g) / float64(255)
	b2 := float64(b) / float64(255)

	// instantiate values for min and max
	min := float64(1)
	max := float64(0)

	// get max and min for the chroma
	colorArray := []float64{r2, g2, b2}

	// loop through array and determine min an max
	for _, color := range colorArray {
		if color > max {
			max = color
		}

		if color < min {
			min = color
		}
	}

	fmt.Println("max: ", max, "min: ", min)
	// HSV values of hex color
	chr := max - min
	hue := float64(0)
	val := max
	sat := float64(0)

	if val > 0 {
		// Calculate Saturation only if Value isn't 0
		sat = chr / val
		if sat > 0 {
			if r2 == max {
				hue = 60 * (((g2 - min) - (b2 - min)) / chr)
				if hue < 0 {
					hue += 360
				}
			} else if g2 == max {
				hue = 120 + 60*(((b2-min)-(r2-min))/chr)
			} else if b2 == max {
				hue = 240 + 60*(((r2-min)-(g2-min))/chr)
			}
		}
	}

	// set the values to the map to return
	m._id = color._id
	m.name = color.name
	m.color = color.color
	m.blue = int16(b)
	m.red = int16(r)
	m.green = int16(g)
	m.chroma = chr
	m.hue = hue
	m.luma = 0.3*r2 + 0.59*g2 + 0.11*b2
	m.sat = sat
	m.val = val

	return m
}

// function to sort colors
func createSortableArray(array []Color) *[]ColorObject {
	finalArray := new([]ColorObject)

	for _, color := range array {
		createColor(color)
	}

	return finalArray
}

func main() {
	colorArray := []Color{
		{
			color: "#EFDECD",
			name:  "Almond",
			_id:   "6057bdf81656909f200a28ed",
		},
		{
			color: "#CD9575",
			name:  "Antique Brass",
			_id:   "6057bdf81656909f200a28ee",
		},
		{
			color: "#78DBE2",
			name:  "Aquamarine",
			_id:   "6057bdf81656909f200a28f1",
		},
		{
			color: "#87A96B",
			name:  "Asparagus",
			_id:   "6057bdf81656909f200a28ef",
		},
	}

	// should be this
	// blue: 117
	// chroma: 0.3450980392156863
	// color: "#CD9575"
	// green: 149
	// hue: 21.818181818181827
	// luma: 0.6363921568627452
	// name: "Antique Brass"
	// red: 205
	// sat: 0.4292682926829269
	// val: 0.803921568627451
	// _id: "6057bdf81656909f200a28ee"

	// blue: 205
	// red: 239
	// green: 222
	// chroma: 0.1333333333333333
	// color: "#EFDECD"
	// hue: 30
	// luma: 0.8832549019607843
	// name: "Almond"
	// sat: 0.14225941422594138
	// val: 0.9372549019607843
	// _id: "6057bdf81656909f200a28ed"

	// blue: 107
	// chroma: 0.24313725490196075
	// color: "#87A96B"
	// green: 169
	// hue: 92.90322580645162
	// luma: 0.596
	// name: "Asparagus"
	// red: 135
	// sat: 0.3668639053254438
	// val: 0.6627450980392157
	// _id: "6057bdf81656909f200a28ef"

	// blue: 226
	// chroma: 0.4156862745098039
	// color: "#78DBE2"
	// green: 219
	// hue: 183.96226415094338
	// luma: 0.7453725490196077
	// name: "Aquamarine"
	// red: 120
	// sat: 0.4690265486725664
	// val: 0.8862745098039215
	// _id: "6057bdf81656909f200a28f1"

	// call the function to create the new objects
	fmt.Println(createSortableArray(colorArray))
}
