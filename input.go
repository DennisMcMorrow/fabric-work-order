package main

import "fmt"

func getTableDimensions(table Table) Table {
	table.customerLength = getLength()
	table.customerWidth = getWidth()
	table.customerDrop = getDrop()
	table.metric = getMetric()
	table.materialName = getMaterialName()
	table.hem = getHem()
	table.surge = getSurge()

	return table
}

func getLength() float64 {
	var customerLength float64
	fmt.Print("Enter length: ")
	fmt.Scanln(&customerLength)
	return customerLength
}

func getWidth() float64 {
	var customerWidth float64
	fmt.Print("Enter width: ")
	fmt.Scanln(&customerWidth)
	return customerWidth
}

func getDrop() float64 {
	var customerDrop float64
	fmt.Print("Enter drop: ")
	fmt.Scanln(&customerDrop)
	return customerDrop
}

func getMaterialName() string {
	var materialName string
	fmt.Print("Enter Material Name (Velvet, Silk, Cotton, Linen, Polyester, etc): ")
	fmt.Scanln(&materialName)
	return materialName
}

func getMetric() string {
	var metric string
	fmt.Print("Enter metric (in, ft, m, cm, mm): ")
	fmt.Scanln(&metric)
	return metric
}

func getHem() bool {
	var hem bool
	fmt.Print("Would customer like fabric hemmed? (True/False): ")
	fmt.Scanln(&hem)
	return hem
}

func getSurge() bool {
	var surge bool
	fmt.Print("Would customer like fabric surged? (True/False): ")
	fmt.Scanln(&surge)
	return surge
}
