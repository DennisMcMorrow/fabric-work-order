package main

import "fmt"

const inchesConversionValue = 0.0277778
const feetConversionValue = 0.333333
const metersConversionValue = 1.09361
const centimetersConversionValue = 0.0277778
const millimetersConversionValue = 0.00109361

func metricToYards(table Table) WorkOrder {
	unitsUsed := getMetricValue(table.metric)
	rawYardDrop := table.customerDrop * unitsUsed
	hem := checkIfHemNeeded(table.hem)
	surge := checkIfSurgeNeeded(table.surge)
	tableFabricLengthInches := table.customerLength + (table.customerDrop * 2)

	workOrder := WorkOrder{
		rawYardLength: ((table.customerLength * unitsUsed) + hem + surge) + (rawYardDrop * 2),
		rawYardWidth:  (table.customerWidth * unitsUsed) + (rawYardDrop * 2),
		rawYardDrop:   rawYardDrop,
	}

	fmt.Println(tableFabricLengthInches)
	fmt.Println(getRollWidth(table.materialName))

	if checkRotate(tableFabricLengthInches, getRollWidth(table.materialName)) { // if the end fabric length is equal to the width of the fabric roll width then swap length and width values
		temp := workOrder.rawYardLength
		workOrder.rawYardLength = workOrder.rawYardWidth
		workOrder.rawYardWidth = temp
		fmt.Println("rotate found")
	}

	return workOrder
}

func getMetricValue(metric string) float64 {
	var unitsUsed float64
	if metric == "yd" || metric == "yards" {
		return unitsUsed
	} else if metric == "in" || metric == "inches" {
		unitsUsed = inchesConversionValue
	} else if metric == "ft" || metric == "feet" {
		unitsUsed = feetConversionValue
	} else if metric == "m" || metric == "meters" {
		unitsUsed = metersConversionValue
	} else if metric == "cm" || metric == "centimeters" {
		unitsUsed = centimetersConversionValue
	} else if metric == "mm" || metric == "millimeters" {
		unitsUsed = millimetersConversionValue
	} else {
		fmt.Println("Metric not defined")
	}

	return unitsUsed
}

func getRollWidth(material string) float64 {
	if roll, ok := FabricRolls[material]; ok {
		return roll.Width
	} else {
		fmt.Println("Material not found:", material)
	}
	return 0
}

func convertYardageToFraction(workOrder WorkOrder) FractionWorkOrder {
	length := int(workOrder.rawYardLength)
	lengthDecimalPart := workOrder.rawYardLength - float64(length)

	width := int(workOrder.rawYardWidth)
	widthDecimalPart := workOrder.rawYardWidth - float64(width)

	lengthFraction := getLengthFraction(lengthDecimalPart)
	widthFraction := getWidthFraction(widthDecimalPart)

	if lengthFraction == " +1" {
		length++
		lengthFraction = ""
	}

	if widthFraction == " +1" {
		width++
		widthFraction = ""
	}

	fractionWorkOrder := FractionWorkOrder{
		formattedLength: fmt.Sprintf("%d%s", length, lengthFraction),
		formattedWidth:  fmt.Sprintf("%d%s", width, widthFraction),
	}

	return fractionWorkOrder
}

func getLengthFraction(decimalPart float64) string {
	switch {
	case decimalPart <= 0.001:
		return " "
	case decimalPart <= 0.125:
		return " 1/8"
	case decimalPart > 0.125 && decimalPart <= 0.25:
		return " 1/4"
	case decimalPart <= 0.33:
		return " 1/3"
	case decimalPart <= 0.375:
		return " 3/8"
	case decimalPart <= 0.5:
		return " 1/2"
	case decimalPart <= 0.625:
		return " 5/8"
	case decimalPart <= 0.67:
		return " 2/3"
	case decimalPart <= 0.75:
		return " 3/4"
	case decimalPart > 0.75 && decimalPart <= 0.875:
		return " 7/8"
	default:
		return " +1" // Indicate rounding up to the next whole number
	}
}

func getWidthFraction(decimalPart float64) string {
	return getLengthFraction(decimalPart)
}
