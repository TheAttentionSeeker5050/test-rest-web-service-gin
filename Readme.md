# My First Golang Gin Server

# Description

This is just a basic project to get familiarized with golang, I just got started with go and want to get familiarized with the syntax, golang's methodology and the processes, methods and capabilities of Gin before I do more complex projects. The project consist on a Rest API in which you make a request and it calculates a specific mathematical function and returns values based on your input. I may try out other features of gin such as adding web drivers, authentication, permissions and advanced validation, but for now I will just make  API post requests without any auth and get the result on a json response

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


### Request Body Structure for login calls

### Request Body Structure for register calls


# Folder structure of the project

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


