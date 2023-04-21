# Shipment

## Application is writen on Golang.

 #### The application is able to do 3 things:

- Get a list of all shipments that have been sent to the system.
- Add a new shipment to the system.
- Get a single shipment by it's ID.

 ### There are 3 endpoints in the application:
--------
- **GET** - localhost:8080/api/shipment (_get a list of all shipments that have been sent to the system_)
#### Response (example):
  ```sh
{
    "message": "Shipments received!",
    "shipments": [
        {
            "id": 1,
            "fromName": "Tom",
            "fromEmail": "tomtop265@gmail.com",
            "fromAddress": "Volrat Thamsgatan 4, Guteborg 41260",
            "fromCountryCode": "SE",
            "toName": "Alex",
            "toEmail": "super12@gmail.com",
            "toAddress": "Broadway 122, New York 13337",
            "toCountryCode": "US",
            "weight": 65,
            "price": 2000
        },
        {
            "id": 2,
            "fromName": "Vital",
            "fromEmail": "vitaltop26@gmail.com",
            "fromAddress": "Ciasna 45, Warshaw 45267",
            "fromCountryCode": "PL",
            "toName": "Sergy",
            "toEmail": "sergy322@gmail.com",
            "toAddress": "Broadway 152, New York 13337",
            "toCountryCode": "US",
            "weight": 21.5,
            "price": 450
        }
    ]
}
```
--------
- **GET** -  localhost:8080/api/shipment/:id (_get a single shipment by it's ID_)
#### Response (example):
  ```sh
{
    "message": "Shipment is getted!",
    "shipment": {
        "id": 2,
        "fromName": "Vital",
        "fromEmail": "vitaltop26@gmail.com",
        "fromAddress": "Ciasna 45, Warshaw 45267",
        "fromCountryCode": "PL",
        "toName": "Sergy",
        "toEmail": "sergy322@gmail.com",
        "toAddress": "Broadway 152, New York 13337",
        "toCountryCode": "US",
        "weight": 21.5,
        "price": 450
    }
}
```
--------
- **POST** -  localhost:8080/api/shipment (_add a new shipment to the system_)
#### Request (example):
```sh
{
    "fromName": "Tom",
    "fromEmail": "tomtop265@gmail.com",
    "fromAddress": "Volrat Thamsgatan 4, Guteborg 41260",
    "fromCountryCode": "SE",
    "toName": "Alex",
    "toEmail": "super12@gmail.com",
    "toAddress": "Broadway 122, New York 13337",
    "toCountryCode": "US",
    "weight": 65
}
```
  #### Response (example):
```sh
{
    "message": "Shipment is added!",
    "price": 2000
}
```
--------
 ### Start the application:

1. Clone or download this repository (https://github.com/Taras-Rm/shipment).
2. Install **Postgress** and **Go** on computer.
3. Enter into the project folder.
4. Create **.env** file and add relevant data to this file (example):
+ SERVER_PORT=:8080
+ DB_NAME=postgres
+ DB_PASSWORD=postgres
+ DB_HOST=localhost
+ DB_USER=postgres
5. Run the application (**go run main.go**).
