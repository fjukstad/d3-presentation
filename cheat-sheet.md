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

# 3 

Putting it all togheter

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
