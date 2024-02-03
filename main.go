package main

import "fmt"

func main() {
	var table Table
	var workOrder WorkOrder
	var fractionWorkOrder FractionWorkOrder
	var customerWorkOrder CustomerWorkOrder

	fmt.Println("Fabric Calculation Tool")

	table = getTableDimensions(table)
	printTableDimensions(table)

	workOrder = metricToYards(table)
	printWorkOrder(workOrder)

	fractionWorkOrder = convertYardageToFraction(workOrder)
	printFractionWorkOrder(fractionWorkOrder)

	customerWorkOrder = CustomerWorkOrder{
		rawYardLength: workOrder.rawYardLength,
		rawYardWidth:  workOrder.rawYardWidth,
		rawYardDrop:   workOrder.rawYardDrop,
	}
	printCustomerWorkOrder(customerWorkOrder)
}
