let incrementButton = document.getElementById("increment");
let decrementButton = document.getElementById("decrement");
let counter = document.getElementById("counter");
let count = 0;

incrementButton.addEventListener("click", function () {
  count++;
  counter.textContent = count;
  updateStyles(count);
});

decrementButton.addEventListener("click", () => {
  count--;
  counter.textContent = count;
  updateStyles(count);
});

function updateStyles(count) {
  if (count < 0) {
    counter.classList.add("negative");
  } else if (count > 0) {
    counter.classList.add("positive");
  } else {
    counter.className = "";
  }
}
