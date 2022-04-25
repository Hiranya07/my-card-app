# my-card-app
## Description
This application contains business logic for accounts/transaction management of the card issue by any financial institution

### How to use(with out docker)
1. Clone this repo using below command: git clone git@github.com:Hiranya07/my-card-app.git
2. cd my-card-app
3. Refer to the configuration section below to update the required configurations
4. Run go mod tidy
5. execute go run main.go

### Configurations
 - This project use mysql as DB for storing data.The Database name,Database address & Database password are configurable as per the need.
 - Before running the application the use need to visit config.json found under my-card-app directory
 - User need to provide value for DbName,DbAddr & DbPassWord configurations
 [ config.json ] https://github.com/Hiranya07/my-card-app/blob/main/config.json

### Rest Endpoints
[Restendpoints] :https://github.com/Hiranya07/my-card-app/blob/main/app/swagger.yml


### Repo contact persons:
- hiranyadeka07@gmail.com

### Used infrastructure components
- mysql

### pacakges
- accounts: Contains business logic to create/retrieve an account
- accounts/db/repositories : Contains logic for accounts related crud operations
- app : Contains the swagger file
- config : Contains the configuration object
- db : Contains the logic for one time database connection creation logic
- models: Contains request/response models
- response: Contains logic for formating the response to be sent for rest api
- routes : Contains the API endpoints handler functions
- transactions: Contains the logic for transaction creation/retrieval
- transaction/db/respositories : Contains logic for transaction related crud operations

### mocking
- Execute the below command to generate mock for the interaces: mockgen -source=path_to_interface destination_path








