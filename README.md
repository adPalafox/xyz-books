# XYZ Books

## About the program
This web app XYZ Books is an inventory system for for books. The backend built with Go (golang) with Vue as the frontend.

The backend is setup in such a way that will make future additions to the functinality of the system easier. Clean Architecture is applied to decouple layers with interfaces. Passing of data between layers are handled by Entities, DTO, DAO, and the backend handles it's own validations.

SQLite is used as file based storage for the book information. While gin is used for handling context and routing.

### Flow
The procees is started at the frontend with user interaction that triggers the request to the backend. The backend exposed routes handles the request and filters them first at the Controller layer.

After input validation, the data is passed to the useccase where it will be further processed. If there needs to have a database interaction, it calls the repository layer to fetch/update the data. Then the presentation layer will return the result back to the caller as response. Throughout all these procedure, the Entity, DTO, and DAO are utilized.

## Run the program
To run the program you would need to have Go (golang) installed in your local machine. 

### Load the backend
First time running the program with migration seeds:
```
make run-migrate
```
To run the program:
```
make run-local
```
Please make sure that the port 8080 is available
### Build and run the frontend
You can run the frontend with the following:
```
cd frontend
npm i
npm run dev
```
Similarly, please ensure that the port 5173 is available, otherwise the frontend will not be able to make requests to the backend even if the frontend server will run on another port in case it is occupied.
Port 5173 is used by the frontend local server while port 8080 should be occupied by the backend.

Also please make sure to perform the commands under the project directory
