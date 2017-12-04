## Our first route
Go comes with a built-in web server and multiplexer. A web server listens for incoming requests, and handles outgoing requests. 

A multplexer, or mux for short, routes incoming request paths to functions. 

## Using templates
We took the first steps towrds building a miniature library application by crating a basic web server. If you want to build a library app, were going to need an easy way to get data in and out of hte system. In this vedio, we'll learn how to use a templating engine to bjild a data-driven user interface. Initially, we'll modify our original route to render a template. We'll send the server query parameters and learn how to access those parameters from Go. Then, we'll use the data from the query parameter to dynamically change our website's content. 

Let's quickly get reintroduced to our current web server. Inside our main function, we created a route handler for the route of our web server, which responds to the request with a short greeting. We also stated our wb erver listening on Port 8080. You can see the result when you visit the page at loalhost:8080 in your browser. Staqtic text generated in code is geat for a start. But let's change the way twe think about rendering the user interface. We can use the concept of the template to display HTML to our users.

This page contains all teh documentation for the standard library of html/template. And you should remember this page for reference later. Let's start by creating an HTML file called index.html in a directory called templates. We'll add some standard HTML boilerplate. A primary HTML node containing empty head and body tags. Insdie the body tag, let's add some text so we can identify our template. Now, we need to update our main dot Go file.

We will start by importing the html/template library alongside our current import statements. Next, we use the template library to parse our templaet. Above the route handler, we will initialize the tempalte's variable with teh call to tempalte Parsefiles. Parsefiles will build tempalte objects from a list of filenames and also return an error.  In our caes, we will enter tempalte/index html to specify our file. We've also wrapped this in a call to tempalte Must, which will absorb the error from Parsefiles and halt execution of the program if it can't parse the template. 

Let's modify our route handler to execute the template. We'll call execute template on a template's object with the response writer object. The name of our index template, nil, as the third parameter for th time being. ExecutiveTemplate returns an error object. We should check to see if this error is nil. If it isn't, we will use http.Error to alert the user that we've encountered a problem. Http.Error takes the response writer, the error text and an error code.  

