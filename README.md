## Our first route
Go comes with a built-in web server and multiplexer. A web server listens for incoming requests, and handles outgoing requests. 

A multplexer, or mux for short, routes incoming request paths to functions. 

## Using templates
We took the first steps towrds building a miniature library application by crating a basic web server. If you want to build a library app, were going to need an easy way to get data in and out of hte system. In this vedio, we'll learn how to use a templating engine to bjild a data-driven user interface. Initially, we'll modify our original route to render a template. We'll send the server query parameters and learn how to access those parameters from Go. Then, we'll use the data from the query parameter to dynamically change our website's content. 

Let's quickly get reintroduced to our current web server. Inside our main function, we created a route handler for the route of our web server, which responds to the request with a short greeting. We also stated our wb erver listening on Port 8080. You can see the result when you visit the page at loalhost:8080 in your browser. Staqtic text generated in code is geat for a start. But let's change the way twe think about rendering the user interface. We can use the concept of the template to display HTML to our users.

This page contains all teh documentation for the standard library of html/template. And you should remember this page for reference later. Let's start by creating an HTML file called index.html in a directory called templates. We'll add some standard HTML boilerplate. A primary HTML node containing empty head and body tags. Insdie the body tag, let's add some text so we can identify our template. Now, we need to update our main dot Go file.

We will start by importing the html/template library alongside our current import statements. Next, we use the template library to parse our templaet. Above the route handler, we will initialize the tempalte's variable with teh call to tempalte Parsefiles. Parsefiles will build tempalte objects from a list of filenames and also return an error.  In our caes, we will enter tempalte/index html to specify our file. We've also wrapped this in a call to tempalte Must, which will absorb the error from Parsefiles and halt execution of the program if it can't parse the template. 

Let's modify our route handler to execute the template. We'll call execute template on a template's object with the response writer object. The name of our index template, nil, as the third parameter for th time being. ExecutiveTemplate returns an error object. We should check to see if this error is nil. If it isn't, we will use http.Error to alert the user that we've encountered a problem. Http.Error takes the response writer, the error text and an error code.  

## Database connections
We created our first route, and we learned how to use a templating engine to render a simple, data driven user interface. For our library app to work we will need to store and access users and their book collections. In this vedio, we'll learn how to connect to a SQLite database from Go, so we can create, read, update, and delete application data. We will start by setting up a SQLite database on our local machine. We'll update our template to display the current database conneciton status

Add some logic to our read handler so we can display to connection status with index.html.  

# Section Two
## Talking to the server
During section one, we built a bare-bones web server in Go. Complete with tempale rendring and a database connection. In section two we'll build on those concepts and start focusing on building features to manage personal book coolections. We'll start this section by sending asynchronous communications from the web-slient to our server. Next we will fetch real library data from an external web-service. Then, we will start crating personal libraries and storing them in the database. Finally we will end this section with an introduction to web middleware to refactor some of our common code.

We will start building a simple search UI suign JavaScript to send request to our web-server. In the new world of the web, it is common practice to load a website and then use javaScript to perform more actions based on user input. Our library applicaiton will be no different. 

If we want to send a request to the Classify API, we will need to make use of hte Get method from the HTTP package. Get takes only a URL and returns a response and an error. We will use the main classify endpoint as our URL, and request books where the title matches our query string. If the error returned from Get is not nil, we will retrun that error along with an empty slice of search results. When we add the search query to the end of our string, we need to make sure to encode the values for the HTTP request. 

Add net/url to the list of the import statements at the top of the file. Back in the search function, we will use the query escape method to escape the input string. Escaping the query will help ensure we have a valid HTTP URL. In order to read the body of the response, we will use ioutil.ReadALL to get the data from our response body, and we will return an error if there is one. Since you want to make sure the body is closed when we are done, we also need to add a defer statement to close the response body at the end of this function. 
## Transaction Flow Basics (Part I)
Step 1: A Client creates and sends a transaction to the Torii gate, which routes the transaction to a peer that is responsible for performing stateless validation. 

Step 2: After the peer performs stateless validation, the transaction is first sent to the ordering gate, which is responsible for choosing the right strategy of connection to the ordering service. 