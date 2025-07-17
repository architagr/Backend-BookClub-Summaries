# Notes for Chapter 1 and Chapter 2

[Link to the recorded session on YouTube](https://youtube.com/live/X_A00q6EaIM)

## Chapter 1 - Introduction to Microservices

### When to use monoliths

1. **Application logic is not defined**; therefore, building a microservice is not advisable.
2. The application has a **minimal scope**.
3. **Low Latency is very important**; hence, we should avoid microservices as they will increase network hops.

### What are some key factors that will drive you to use microservices?

1. A large application is causing **slower deployments**.
2. You can not only deploy just **1 part that gets updated frequently**.
3. **Higher blast radius**, one issue in a library function causes a lot of code to break.
4. **You can not scale components individually**, eg, in an e-commerce application, only the payment system needs to be scaled on the day of sale.
5. **Scaling vertically is also an issue** as now we need not resources on a single server, and increasing that further is hard and expensive.
6. Some part of the app is having more operations, causing slowness in other parts of the application.
7. **Higher security risk**.
8. **Different components of the app need different infra**; eg when we are processing data that the user has requested to be downloaded, or a daily report that needs to be sent
9. **Harder to test**, as some parts need more testing, or may need a performance test, but others do not need that.
10. Different part of the application needs **different tech stacks**.
11. Some parts of the app have **different compliance**, eg, user management needs strict compliance on PII (Personally identifiable information)

But _Software engineering is a result of many **trade-offs**_; there can not be a single rule that will fit perfectly in all cases. Having said that, let’s discuss some pros and cons of using microservices.
These are some things we all should know before deciding between a monolith and a microservice.

### Pros of using Microservices

1. **Fast compilation and build** (per service), and it is also good if we talk about over all application, as all builds can be in parallel.
2. **Fast deployment** - as smaller applications get deployed in parallel.
3. **Different deployment schedules** and monitoring for every application based on requirements.
4. Each component can have **independent testing requirements**.
5. Each microservice is open to **using different languages** based on feasibility.
6. **Horizontal scaling** is easy.
7. **Hardware flexibility**.
8. **Falult Isolation**.
9. **Distributed Development**.
10. **Cost optimisation**.
11. Easy decision-making per microservice.

### Cons of microservices

1. Higher network usage, as microservices have to communicate with each other.
2. Hard debug.
3. Increased dependency on integration tests.
4. Consistency and transaction are hard to maintain as data is distributed across multiple databases.
5. Duplication, overlapping functionality (this can be solved by a Mono repo).

## Chapter 2 - Scaffolding a Go Microservice

Challenges faced by engineers working on large Go applications:

1. Finding the right project structure to make code easy to evolve.
2. Writing idiomatic Go code.
3. Separation of components of microservices.

Let’s start by understanding some common principles to **write idiomatic Go Code**.

### Things I was doing

1. Exporting names that start with Caps or uppercase
2. If another package imports your package and uses the structure, variable or interface, then they will have to prefix that with the package name, so we should not define the name with the package name. eg, Buffer is a struct name, not BytesBuffer or BufferBytes, as other packages using it will use it like bytes.Buffer
3. A setter can have a name starting with `Set`. eg. SetName(string)
4. Package name abbreviations should be avoided if not widely used, like fmt, or cmd
5. Every exported member should have a comment describing it.
6. Comments should end with a period.
7. The first sentence of the comment begins with the member's name.
8. Avoid panics until it makes sense.
9. The error string should start with a lowercase letter.
10. Errors should not end with punctuation marks.
11. Always handle the error as soon as it is returned from a function.
12. Use error wrapping when you want to add more info, with the use of ‘%w’
13. To check the error object type or error content, use the errors package functions ‘Is’ and ‘As’. as = =' operator might cause issues if the function starts adding additional info using the error wrapping.
14. The interface should be defined on the caller side.
15. Tests should always give full information about what went wrong.
16. Only write a test for the public function, and that will also cover private functions.
17. Always pass context as the 1st parameter of the function performing some IO calls.
18. Do not attach context to structures, as this is called context leakage.

### Things that I will start doing

1. The package name should be a single word in lowercase.
2. Getter function should not have names starting with `Get`, eg we should not have GetName(); instead, we should have Name()
3. Single method interface should have a name as method name followed by `er`, eg, Writer interface only has Write function
4. Initialism and acronyms should have consistent case, eg either URL or url; we should not use Url or uRL, similarly ID or id is correct Id is not right.
5. Every package should have a comment describing its contents.
6. Handle each error, do not discard that using the ‘\_’ assignment.

### Scaffolding backend service

In the book author has discussed a movie rating application, but to make sure we have understood what he wants to deliver, let’s consider a different yet simple application.

Let’s design an **inventory management application**, and before considering if it should be a microservice or a monolith, let’s try evaluating all points. But even before we go in that discussion, let’s first understand the application.
We want to design an inventory management application where info that we will have to maintain is:

1. **Metadata** - Basic information of the product, like
   1. manufecturer
   2. specification
   3. type of product (perishable, etc.)
   4. listed price on the product
   5. Category
   6. Sub category
2. **Orders** - all orders we have received for the product will also help with any replenishment or return
3. **Discount** - if we have any offer going on a product or category, but to start with, we will keep this simple and only have a discount applied to a product.

Now the user can ask for a list of products based on category or subcategory, we will also have to tell them the items left in stock and the final price after discount
The user can place a buy/sell/return order. Ideally, we will do validation when the return order is there, and we will not allow buying more than the available stock.

### Let's see if we can use monoliths architecture

1. Application logic is not defined; therefore, building a microservice is not advisable.
   1. It is defined, hence we may or may not choose a monolith
2. The application has a minimal scope
   1. No, scope can grow, as we might start adding discounts on categories, or add verified sellers, add quality checks for sellers, etc.
   2. Hence, we may or may not choose a monolith.
3. Low Latency is very important; hence, we should avoid microservices as they will increase network hops.
   1. We can have no requirement for low latency, as the only service that will need that will be the payment service, which is not what we are planning.
   2. Hence, we may or may not choose a monolith.

### Let's see if we can use microservices architecture

1. A large application is causing slower deployments
   1. This will be true if we talk about a scale equal to the Amazon scale
   2. Definately inclining towards using microservice
2. You can not only deploy just 1 part that gets updated frequently.
   1. This will be important as we might have to do frequent updates on the order app due to government rule changes
   2. Definately inclining towards using microservice
3. Higher blast radius, one issue in a library function causes a lot of code to break.
   1. This can be there as if the order app has an issue and we are on a monolith, then the full website is down, but in a microservice, we can show a list of products atleast
   2. Definately inclining towards using microservice
4. You can not scale components individually, eg, in an e-commerce application, only the payment system needs to be scaled on the day of sale.
   1. Yes, this will be there as order apps data will be huge, and traffic is also expected to grow on days we have sales
   2. Definately inclining towards using microservice
5. Scaling vertically is also an issue as now we need not resources on a single server, and increasing that further is hard and expensive.
   1. If the number of orders grows, we can not keep adding database space for the complete application database.
   2. Definately inclining towards using microservices.
6. Some part of the app is having more operations, causing slowness in other parts of the application.
   1. Yes order will be that component
   2. Definately inclining towards using microservice
7. Higher security risk
   1. The order app will have the maximum security risk as it involves finances
   2. Definately inclining towards using microservice
8. different components of the app need different infra; eg when we are processing data that the user has requested to be downloaded, or a daily report that needs to be sent
   1. The order app will need notification and tracking-related infrastructure
   2. Definately inclining towards using microservice
9. harder to test, as some parts need more testing, or may need a performance test, but others do not need that.
   1. yes
   2. Definately inclining towards using microservice
10. Different part of the application needs different tech stacks.
    1. do not know at this point
    2. no opinion due to this, to choose between microservice and monolith
11. Some parts of the app have different compliance, eg, user management needs strict compliance on PII (Personally identifiable information)
    1. Yes, the financial transition will have this.
    2. Definately inclining towards using microservice.

Note: And as we are discussing microservices, let’s think we have to go for microservices.
