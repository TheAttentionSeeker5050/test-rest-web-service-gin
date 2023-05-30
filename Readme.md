# My First Golang Gin Server

## Description

This is just a basic project to get familiarized with golang, I just got started with go and want to get familiarized with the syntax, golang's methodology and the processes, methods and capabilities of Gin before I do more complex projects. The project consist on a Rest API in which you make a request and it calculates a specific mathematical function and returns values based on your input. I may try out other features of gin such as adding web drivers, authentication, permissions and advanced validation, but for now I will just make  API post requests without any auth and get the result on a json response. 
**Disclaimer:** Please have in mind that  this is a work in progress and not everything below is ready yet

<hr>

## Tools to be Used 
- A server technology, in this case, as I want to tinker with golang, it will be golang gin
- A database, first option is postgresql, second option can be a mongo atlas instance on local computer or docker for development
- Postman to end to end testing of API calls
- Testing library for golang, I will have to research on this
- If I find any tool that can make this work easier on testing, or input validation, I will use it
- Input validation will be regex or basic data type / number validation when Necessary
- For authentication may go for cookies validation for the moment, as it is not that complicated in most languages, IDK how it is on golang but will see

<hr>


## API Structure

### Address Structure

- api/v1
    - /calculator/ (no api calls allowed on this exact address, but on children)
        - all the calculators names are children of this, post request only
            - basic-calculator
            - bin-to-hex-converter
            - hex-to-bin-converter
            - sample-statistics-calculator
        - /history, get request, to get all calculator operations saved by him and other users
        - /history/<some filter> may be possible in the future
    - /user (no api calls allowed on this exact address, but on children)
        - /profile, get request, user has to be authenticated, may do delete and put request for editing profile data in the future, but for now only get
        - /login, post only, if authentication successful user can acces data on its profile
        - /register, post only, user has to /login to authenticate


### Request Body Structure for Calculator calls
- /api/v1/calculator/basic-calc → **post** request; **request-body**: {params: (num1, num2, operand)}; **response**: {responseStatus, requestParams, date, result: (value), user(username or anon user)}
- /api/v1/calculator/hex-to-bin → **post** request; **request-body**: {params: (hexString)}; **response**: {responseStatus, requestParams, date, result: (binString), user}
- /api/v1/calculator/bin-to-hex → **post** request; **request-body**: {params: (binString)}; **response**: {responseStatus, requestParams, date, result: hexString, user}
- /api/v1/calculator/statistics-calc → **post** request; **request-body**: {params: (… nums)}; **response**: {responseStatus, requestParams, date, result: (mean, median, standardDeviation), user}
- /api/v1/calculator/history → **get** request; **response**: {responseStatus, calcType, requestParams, date, result, username}


### Request Body Structure for user calls
- /api/v1/user/profile → **get** request; **response**: {responseStatus, username, firstName, lastName, userSince, emailAddress, birthDay}; **needs authentication**
- /api/v1/user/login → **post** request; **request-body**: {usernameOrEmailAddress, password}; **response**: {responseStatus, responseMessage}; **needs authentication**
- /api/v1/user/register → **post** request; **request-body**: {username, firstName (optional), lastName (optional), emailAddress, password, passwordConfirmation, birthDay (optional)}; **response**: {responseStatus, , responseMessage}; **needs authentication**

<hr>

## Folder structure of the project

- root (package main)
    - routes 
    - common
        - calculator: calculator function directory
        - validator: validator functions directory
        - authentication:authentication methods directory
        - other: stuff such as authentication middleware (may change)
    - config
        - db, environment variables, authentication, testing, url validation, etc configuration files
    - model
        - queries, views, other model files
    - controller
        - the middle point between the api routes and the model back and forth, uses common functions
        - Invalid call api views
    - tests
        - all tests files go here (probably split it in subdirectories by purpose)

    - main file
    - readme file
    - mod and checksum files (.mod and .sum)

<hr>

## Updates

Updates and milestones go here