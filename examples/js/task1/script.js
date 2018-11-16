var trace = {
  x: [],
  y: [],
  mode: 'markers'
};

function rand(min, max){
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

min1 = 1
max1 = 5
min2 = 1
max2 = 7

d = Math.abs(parseInt(max2) - parseInt(max1))

for(i = 0; i < 500; i++) {
   sh1 = rand(parseInt(min1), parseInt(max1))
   sh2 = rand(parseInt(min1), parseInt(max1))

   if(sh2 < d){
     sh1 = max1 + sh2
   }

   trace.x.push(i)
   trace.y.push(sh1)
}

var data = [trace];

var layout = {};

Plotly.newPlot('myDiv', data, layout);
