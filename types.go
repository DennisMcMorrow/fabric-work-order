package main

type Table struct {
	customerLength, customerWidth, customerDrop float64
	metric                                      string
	hem, surge                                  bool
}

type WorkOrder struct {
	rawYardLength, rawYardWidth, rawYardDrop float64
}

type FractionWorkOrder struct {
	formattedLength, formattedWidth string
}

type CustomerWorkOrder struct {
	rawYardLength, rawYardWidth, rawYardDrop float64
}
