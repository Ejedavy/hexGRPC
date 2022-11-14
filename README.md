#  About hexGRPC and project structure
A simple GRPC server in GoLang to perform basic arithmetics. <br/>
Including the MySQL as the database and using SQL scripts to start the docker db service <br/>
Using GRPC as the main endpoint type <br/>

Making the application scalable and maintanable by making the layers low coupled and high cohesion.<br/> The database and endpoint type can easily be replaced without changing the core of the application


# Test
Included unit tests to test the core layer of the application and an end to end test to test the workflow of the system. 

# To Run
You need to run docker compose up to start the services and run the tests
