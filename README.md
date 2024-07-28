# go-order-app

An application built using Go and the Gin framework. This is the first time I have used Golang.

## Rest API

It contains three Rest API endpoints:
### GET /api/packs
This endpoint is to get a JSON response containg the combination of packs required to fulfill an order based on a number of items ordered. It requires a query parameter 'itemsOrdered'. It returns the response as JSON.

### PUT /api/pack-sizes
This endpoint is to change the array of pack sizes used to fulfill the orders. It requires a JSON payload, containing one attribute 'newPackSizes' which is an array of integers. Currently, it only changes the array in memory. Extending this would be to persist the new pack sizes in a database.

### PUT /api/reset-pack-sizes
This endpoint is to reset the array of pack sizes to the default: [250, 500, 1000, 2000, 5000].
It does not require a payload.

## Frontend

The frontend is very basic (very!). 
It has two routes:
### GET /
The index page contains a form that allows you to submit a number of items. This calls the GET /pack-sizes route, passing the number in the form to the route via a query parameter. 

### GET /pack-sizes
This route calculates the packs required, and passes the data back to the index template. It renders a list of the required packs. Its server side validation currently returns JSON as opposed to rendering an error message on the frontend, as the route utilises the same middleware function as the equivalent REST API endpoint.

## Improvements

To improve this application, I would:
- add unit tests for the frontend route
- for a production application, the frontend would be a separate component to the REST API, and the frontend would call the API endpoints to retrieve the data.
- as mentioned previously, persist changes in the pack sizes in a database.
- improve the UI (actually have a UI!)
- structure the application more appropriately.
- utilise interfaces, both to follow SOLID principles as well as allow mocking for the unit tests to only test specific modules. This I would have ideally done but the learning curve for doing this in Golang prevented me from doing it on this application.
- Further jobs on the GitHub workflow: coverage, code quality, vulnerability checks etc.
- Utilised HTMX to pass partial HTML back to template to render (I am still learning HTMX).

## Deployed Application

The application is deployed on Heroku. This is currently a manual process to deploy from the remote main branch on the GitHub repo. For a production application this would be deployed as part of a CI/CD pipeline after a Pull Request has been merged into main.

The url for the deployed app is: https://go-order-app-5a35dfeb6589.herokuapp.com/ 

If the URL doesn't work for whatever reason, please let me know.