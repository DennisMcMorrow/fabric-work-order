# Fabric Work Order Project
A project designed to handle work orders for fabric, providing necessary calculations and data management.


## Features
- Input measurement information into the GUI
- Select radio buttons and selectors to meet customer criteria
- Calculate and provide yardage in work order and customer formats
- Manage fabric data with a DBMS


## Terminologies
- Hem: 1 or 2 inch tuck for fabric stitching
- Surge: 1 or 2 inch tuck for fabric stitching (rarely offered)
- Round: Calculation for a round table (length + (drop x 2))
- Lever: Expandable table
- Double: Use of 2 rolls
- Trim: Applied around round tables, possibly regular tables too
- Square: Type of table
- Rotate: Turning fabric width becomes length


## Technology
- Programming Language: Golang

## Database Management
- Fabric Name: STRING (e.g., Velvet, Silk)
- Roll Size: INT (e.g., 54", 72", 110")
- Amount Remaining: INT (Initial yardage per roll)
- Hem Needed: BOOLEAN (Material dependent)
- Surge Needed: BOOLEAN (Material dependent)


## Program Functionality
- Employees input fabric measurements (length, width, drop)
- GUI calculates necessary yardage for work orders
- Handles measurements in inches, feet, or meters
- Outputs yardage in two formats: work order and customer measurement


For more details, visit the [GitHub repository](https://github.com/DennisMcMorrow/fabric-work-order).
