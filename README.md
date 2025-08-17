![CI status](https://github.com/MartinGallauner/goffeine/actions/workflows/ci.yml/badge.svg)

# Goffeine

## About
Goffeine is a simple caffeine tracking tool.
It allows you to add your caffeine consumption in natural language. 

The half-life of caffeine is set to 5 hours and caffeine intake that's more than 24 hours in the past is ignored.
This is set according to [this paper](https://www.ncbi.nlm.nih.gov/books/NBK223808/#:~:text=The%20mean%20half%2Dlife%20of,et%20al.%2C%201989)
I picked because it's been the very first result of my minimal effort google search.



## Quick Start

1. You can install the dependencies by running `make setup`.
2. Rename the provided `.env.dev` to `.env`, add your own OpenAPI key.
3. To start the project locally run `make run`. This will run the complete build and starts the server


## Run the tests
-  You can start all tests and prepared linters with `make test`.  

## Contributing 

The project is too half-baked to accept contributions.



