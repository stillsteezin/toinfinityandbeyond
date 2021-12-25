const starfield = document.getElementById("starfield");
var numStars = 500;

for($i = 0; $i < numStars; $i++)
{
  var star = document.createElement("div");
  star.classList.add("star");

  var speed = Math.random() * 1000;
  var rotation = Math.random() * 360;
  var delay = Math.random() * 1000;
  var scale = Math.random() * 10;

  star.setAttribute("style","--speed: " + speed + "ms; --rotation: " + rotation + "deg; --delay: " + delay + "ms; --scale: " + scale);
  starfield.appendChild(star);
}