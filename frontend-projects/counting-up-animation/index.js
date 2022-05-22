const counters = document.querySelectorAll(".counter");
const speed = 400; // speed of counting

counters.forEach((counter) => {
  const updateCounter = () => {
    // + symbol converts a string to number.
    const target = +counter.getAttribute("data-target");
    const count = +counter.innerText;

    const increment = target / speed;

    if (count < target) {
      counter.innerText = Math.floor(count + increment);
      setTimeout(updateCounter, 1);
    } else {
      counter.innerText = target;
    }
  };
  updateCounter();
});
