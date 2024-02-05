Fabric Work Order Project

Technologies Used:
Programming Language: Golang
APIS: External API Interaction/RESTful API
Containers: Docker/Kubernetes
DBMS: PostgreSQL
Cloud Service: AWS CLI, AWS Console, S3
Implement: Data Structures, Design Patterns, Network Protocols
CI/CD Environment: Jenkins, GitLab CI
NoSQL/Big Data Technologies: Hadoop, HBase, Cassandra, MongoDB


Program Functionality:
User will be able to measurement information into the GUI and be able to select radio buttons and selectors to meet the criteria of what is being asked for from the customer

For example, customer walks in to get a specific type of fabric and will provide the employee with length, width, and drop they are looking for in the fabric.
Employee will then plug in these values into the GUI and be given the correct calculations needed for the work order.
Customers relay this information most of the time in inches, feet, or even meters. 
We must calculate this information and give the yardage back in two different forms, form1 = work order measurement and form2 = measurements given to the customer

There are certain terminologies needed to be understood-
Hem = 1 OR 2 inch tuck for fabric to be stitched (material dependent) (1/2 inch all around)
Surge = 1 OR 2 inch tuck for fabric to be stitched (material dependent) (not offered most of the time)
Round = table (l + (drop x 2))
Lever = expandable table
Double = means to use 2 rolls
Trim = around a round table possibly applied to regular tables too
Square = type of table
rotate = turn fabric width becomes length

Width = width of the object
Length = length of the object
Drop = length of the drop off the object

We can use a database management system to manage the data for all the fabrics and their certain fields, field examples are this:
FABRIC NAME = STRING (Velvet, silk, etc)
ROLL SIZE = INT (54", 72", 110", etc) (velvet, silk, no pattern then use the length as a width and the width as a length in the work order)
AMOUNT REMAINING = INT (all rolls start at a certain amount of yardage/roll)
HEM NEEDED = BOOLEAN (material dependent, if true then hemp needed else hemp not needed)
SURGE NEEDED = BOOLEAN (material dependent, if true then hemp needed else hemp not needed)



