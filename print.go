package main

import "fmt"

func printTableDimensions(table Table) {
	fmt.Println("TABLE DIMENSIONS")
	fmt.Println("Length: ", table.customerLength)
	fmt.Println("Width: ", table.customerWidth)
	fmt.Println("Drop: ", table.customerDrop)
	fmt.Println("Metric: ", table.metric)
	fmt.Println("Hem: ", table.hem)
	fmt.Println("Surge: ", table.surge)
}

func printWorkOrder(workOrder WorkOrder) {
	fmt.Println("RAW WORK ORDER")
	fmt.Println("Length: ", workOrder.rawYardLength)
	fmt.Println("Width: ", workOrder.rawYardWidth)
	fmt.Println("Drop: ", workOrder.rawYardDrop)
}

func printCustomerWorkOrder(customerWorkOrder CustomerWorkOrder) {
	fmt.Println("CUSTOMER WORK ORDER")
	fmt.Println("Length: ", customerWorkOrder.rawYardLength)
	fmt.Println("Width: ", customerWorkOrder.rawYardWidth)
	fmt.Println("Drop: ", customerWorkOrder.rawYardDrop)
}
