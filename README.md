# nix-trainee-5-6-7-8

🍕Run App with simple command **_make up_** (of course if you have preinstalled Docker).

App started on 8080 port in your machine. After start MySQL container just run **_make init_** for creating Database & structure of Tables

Ok, now we run App, but for use Post Service you should to Register&Login for getting Bearer Access Token. Use this sequence of HTTP requests to achieve it:
1. [POST] http://localhost:8080/v1/users/registration
2. [POST] http://localhost:8080/v1/users/confirmation
3. [POST] http://localhost:8080/v1/users/login
4. Do CRUD operations under Post & Comment objects
5. [POST] http://localhost:8080/v1/tokens/refresh

About most popular command read here - **_make help_**

📚Read & Test with [Swagger Docs](http://localhost:8080/docs/index.html)

🎲Test or Develop with Postman Collection(just import **postman-collection.json** file)

📌I combined exercises 5-6-7-8 into the Post Service. All environment variables(**_.env.dev_** file with passwords & API keys) are intentionally left in the root directory for easy application startup.

❗️Methods registration/confirmation/login/refresh return User data and low level errors for clearly demonstration. All CRUD are covered by tests with SQL mocks.
___
#### JSON
If you want to send JSON data - you must use Content-Type Header with value application/json.
In all cases you will receive data in JSON format except value Header "Access: application/xml" - in this case you will receive XML data
___
#### XML
If you want to send XML data - you must use Content-Type Header with value application/xml.
If you want to receive XML data - you must use Access Header with value application/xml.