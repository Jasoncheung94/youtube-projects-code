function displayDateTime() {
  var date = new Date();
  var hours = date.getHours();
  var minutes = date.getMinutes();
  var seconds = date.getSeconds();
  var meridiemoffset = "AM";

  // handle mightight use case.
  if (hours == 0) {
    hours = 12;
  }

  if (hours > 12) {
    hours = hours - 12;
    meridiemoffset = "PM";
  }

  hours = hours < 10 ? "0" + hours : hours;
  minutes = minutes < 10 ? "0" + minutes : minutes;
  seconds = seconds < 10 ? "0" + seconds : seconds;

  // display the time
  var time = hours + ":" + minutes + ":" + seconds + " " + meridiemoffset;
  document.getElementById("digitalClock").innerText = time;

  // display the date
  let options = {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
  };
  document.getElementById("date").innerHTML = date.toLocaleString(
    "en-GB",
    options
  );
  // timeout of 500ms for accurate time display.
  setTimeout(displayDateTime, 500);
}
displayDateTime();

// optional to have an onload event to display the time.
// window.onload = (event) => {
//     displayDateTime();
// };
