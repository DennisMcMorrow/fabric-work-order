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

	workOrder := WorkOrder{
		rawYardLength: ((table.customerLength * unitsUsed) + hem + surge) + (rawYardDrop * 2),
		rawYardWidth:  ((table.customerWidth * unitsUsed) + hem + surge) + (rawYardDrop * 2),
		rawYardDrop:   rawYardDrop,
	}

	return workOrder
}

func getMetricValue(metric string) float64 {
	var unitsUsed float64
	if metric == "in" {
		unitsUsed = inchesConversionValue
	} else if metric == "ft" {
		unitsUsed = feetConversionValue
	} else if metric == "m" {
		unitsUsed = metersConversionValue
	} else if metric == "cm" {
		unitsUsed = centimetersConversionValue
	} else if metric == "mm" {
		unitsUsed = millimetersConversionValue
	} else {
		fmt.Println("Metric not defined")
	}

	return unitsUsed
}

func convertYardageToFraction(workOrder WorkOrder) FractionWorkOrder {
	length := int(workOrder.rawYardLength)
	lengthDecimalPart := workOrder.rawYardLength - float64(length)

	width := int(workOrder.rawYardWidth)
	widthDecimalPart := workOrder.rawYardWidth - float64(width)

	lengthFraction := getLengthFraction(lengthDecimalPart)
	widthFraction := getWidthFraction(widthDecimalPart)

	fractionWorkOrder := FractionWorkOrder{
		formattedLength: fmt.Sprintf("%d%s", length, lengthFraction),
		formattedWidth:  fmt.Sprintf("%d%s", width, widthFraction),
	}

	return fractionWorkOrder
}

func getLengthFraction(decimalPart float64) string {
	switch {
	case decimalPart <= 0.125:
		return " 1/8"
	case decimalPart <= 0.17:
		return " 1/6"
	case decimalPart <= 0.25:
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
	case decimalPart <= 0.83:
		return " 5/6"
	case decimalPart <= 0.875:
		return " 7/8"
	default:
		return " +1" // Indicate rounding up to the next whole number
	}
}

func getWidthFraction(decimalPart float64) string {
	return getLengthFraction(decimalPart)
}
