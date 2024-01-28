package main

import "fmt"

func main() {
	var table Table
	var workOrder WorkOrder
	var customerWorkOrder CustomerWorkOrder

	fmt.Println("Fabric Calculation Tool")

	table = getTableDimensions(table)
	printTableDimensions(table)

	workOrder = metricToYards(table)
	printWorkOrder(workOrder)

	customerWorkOrder = CustomerWorkOrder{
		rawYardLength: workOrder.rawYardLength,
		rawYardWidth:  workOrder.rawYardWidth,
		rawYardDrop:   workOrder.rawYardDrop,
	}
	printCustomerWorkOrder(customerWorkOrder)
}
