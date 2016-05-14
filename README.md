# RESTumé
Simple little RESTful API to serve my CV.
Heavily influenced by and based upon [Matt Heath's Golang UK Conference talk](https://www.youtube.com/watch?v=cFJkLfujOts).

## The vision
A REST API that exposes my CV for the world to see.
The service should be built in a way that allows more features or services to be quickly added on later.
The code should have good test coverage, with tests not being treated as a second class citizen.

## TODO
- Add better tests
- Build a proper service testing rig, to avoid service config setup in the tests.
- Expand the remaining endpoints out to use the database and not hardcoded data.
- Add authentication to POST requests.
- Consider how to deploy the code. Maybe using google app engine.
- Clean up service.go and split out the handlers code
- Refactor the endpoints to use /{collection}/{index}/{query} as there will be a lot of duplicate code.