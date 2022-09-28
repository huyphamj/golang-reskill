# Golang reskill project

This is an example and `simple` project to mock flight-management system.

It runs under microservices architecture with 3 services:

- Booking: Manage reserved ticket for users and flights.
- User: Manage users.
- Flight: Manage flights

## Project structure

### Common

- Configuration: manage configuration
- Database: contains database connection initialization
- Proto: contains `proto` files for gRPC
- Util: contains util functions

### User

- Located in folder `/user`
- Restful API served:

POST `/user` - register new user

GET `/user/:username` - get user data by username

PUT `/user/:username` - update user data

POST `/user/validatePassword` - validate username/password

- gRPC served:

`UserService`.`ValidatePassword` to validate username/password

### Flight

- Located in folder `/flight`
- Restful API served:

GET `/flight` - Get list flights

GET `/flight/:code` - Get flight by code

- gRPC served:

`FlightService`.`SearchFlight` to query flight data (code, seats, time, ...)

### Booking

- Located in folder `booking`
- Restful API served:

POST `/booking` - Reserve tickets for a flight with amount of tickets desired

GET `/booking` - Get list of user's reserved bookings

GET `/booking/:id` - Get detail of a reservation

## Usage

- This project's purpose is for user viewpoint. Therefore, their actions should be:
- View flights
- View personal data
- Reserve flights

