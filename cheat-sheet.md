# 1
First append to body 
```javascript
var div = document.createElement("div");
div.innerHTML = "Hello, world!";
document.body.appendChild(div);
```


Then let's create a circle. 
```javascript
div.innerHTML = '<svg width="100" height="100"> <circle cx="30" cy="60" r="10"></circle></svg>' 
```


# 2

Using d3 to do the same thing as we did. With d3 we can handle groups of related
elements called *selections*. 
```javascript
var body = d3.select("body");
var div = body.append("div");
div.html("Hello, world!");
```

Another nice thing about selections is *method chaining*
```javascript
var body = d3.select("body").append("div").html("Hello, mate!");
```

Now, let's make some visual stuff! 

# 3


Select them all
```javascript
var circle = d3.selectAll("circle");
```

New color and different radius 

```javascript
circle.style("fill", "steelblue");
circle.attr("r", 30);
``` 

Dancing circles. 

```javascript
circle.attr("cx", function() { return Math.random() * 720; });
```

We need data! Refresh 2.html. 

```javascript
var circle = d3.selectAll("circle");
circle.data([20, 40, 70])
```

Change their radius according to the data we entered. 
```javascript
circle.attr("r", function(d){return d;});
```

Move them a bit from each other, note we're using the index. 
```javascript
circle.attr("cx", function(d,i){return i * 100 + 30;})
```

What if we had more than 3 data elements? We need a new circle! 

```javascript 
var svg = d3.select("svg");

var circle = svg.selectAll("circle").data([20, 40, 70, 90]);

var circleEnter = circle.enter().append("circle");

```

Update the attributes, so that the new one will have them as well 
```javascript
circle.attr("r", function(d){return d/2;});
circle.attr("cy", 60) 
circle.attr("cx", function(d,i){return i * 100 + 30;});
```

Now removing the last two items. 
```javascript
var circle = svg.selectAll("circle").data([20,40]);
circle.exit().remove() 
```

# 4 

Putting it all together

```javascript 

var svg = d3.select("svg") 

var circle = svg.selectAll("circle")
    .data([20,40,70,90], function(d) { return d; });

circle.enter().append("circle")
    .attr("r", function(d){return d/2;})
    .attr("cy", 60) 
    .attr("cx", function(d,i){return i * 100 + 30;});

circle.exit().remove();
```

# 5 Lets make a scatterplot! 

First some data
```javascript
    var dataset = [{x:10, y:10}, 
                    {x:20, y:80}]
``` 

Make a *svg* thing where we can plot and that 
```javascript 
    var svg = d3.select("body")
                .append("svg")
                .attr("width", "300")
                .attr("height", "200"); 
```

Now dataset! 
```javascript 
    var circle = svg.selectAll("circle")
                    .data(dataset) 
```

Surplus elements are removed since they end up in the exit selection
```javascript
    circle.exit().remove() 
```

New data and circles are added 
```javascript
    circle.enter().append("circle")   
          .attr("r", "3");
```

Update the location of existing circles 
```javascript
    circle.attr("cx", function(d) {
                return d.x
           })
          .attr("cy", function(d) {
                return d.y
          });
```

What about axes? 
