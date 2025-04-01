//About this Project

This API works only in case of wrong prefix data and FilmTypeID has start with "FG"

//Installation
//echo framework
1. Import echo to your project
    "github.com/labstack/echo/v4" 
2. run command : go mod init "nameproject"
3. run command : go mod tidy

//Launch Instructions
To start the API use command:
go run ./main.go

//Endpoint Instructions

POST  /order  ->  Send the InputOrder

Output as Json

Example:

 POST  localhost:3000/order


Input:
[
  { 
  "no": 1,
  "platformProductId": "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2", 
  "qty": 1, 
  "unitPrice": 160, 
  "totalPrice": 160 
  }, 
  { 
  "no": 2, 
  "platformProductId": "FG0A-PRIVACY-IPHONE16PROMAX", 
  "qty": 1, 
  "unitPrice": 50, 
  "totalPrice": 50 
  } 
]

Output:
 [ 
{ 
"no": 1, 
"productId": "FG0A-CLEAR-OPPOA3", 
"materialId": "FG0A-CLEAR", 
"modelId": "OPPOA3", 
"qty": 2, 
"unitPrice": 40.00, 
"totalPrice": 80.00 
}, 
{ 
"no": 2, 
"productId": "FG0A-MATTE-OPPOA3", 
"materialId": "FG0A-MATTE", 
"modelId": "OPPOA3", 
"qty": 2, 
"unitPrice": 40.00, 
"totalPrice": 80.00 
}, 
{ 
"no": 3, 
"productId": "FG0A-PRIVACY-IPHONE16PROMAX", 
"materialId": "FG0A-PRIVACY", 
"modelId": "IPHONE16PROMAX", 
"qty": 1, 
"unitPrice": 50.00, 
"totalPrice": 50.00 
}, 
{ 
"no": 4, 
"productId": "WIPING-CLOTH", 
"qty": 5, 
"unitPrice": 0.00, 
"totalPrice": 0.00 
}, 
{ 
"no": 5, 
"productId": "CLEAR-CLEANNER", 
"qty": 2, 
"unitPrice": 0.00, 
"totalPrice": 0.00 
}, 
{ 
"no": 6, 
"productId": "MATTE-CLEANNER", 
"qty": 2, 
"unitPrice": 0.00, 
"totalPrice": 0.00 
}, 
{ 
"no": 7, 
"productId": "PRIVACY-CLEANNER", 
"qty": 1, 
"unitPrice": 0.00, 
"totalPrice": 0.00 
} 
]

//Contact Thanawat Kittisupjaroen 

Email : tkittisupjaroen@gmail.com

