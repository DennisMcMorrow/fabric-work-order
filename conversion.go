package main

import "fmt"

const inchesConversionValue = 0.0277778
const feetConversionValue = 0.333333
const metersConversionValue = 1.09361
const centimetersConversionValue = 0.0277778
const millimetersConversionValue = 0.00109361

func metricToYards(table Table) WorkOrder {
	unitsUsed := getMetricValue(table.metric)

	workOrder := WorkOrder{
		rawYardLength: table.customerLength * unitsUsed,
		rawYardWidth:  table.customerWidth * unitsUsed,
		rawYardDrop:   table.customerDrop * unitsUsed,
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

func getHemDimensions() {

}

func getSurgeDimensions() {

}
