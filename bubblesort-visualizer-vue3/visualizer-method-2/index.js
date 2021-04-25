const app = Vue.createApp({
  data() {
    return {
      array: [],
      size: 30,
    };
  },
  methods: {
    fillArray() {
      // populate array with random numbers between 5 and 750
      this.array = [];
      for (let i = 0; i < this.size; i++) {
        this.array.push(this.getRndInteger(5, 750));
      }
    },
    getRndInteger(min, max) {
      return Math.floor(Math.random() * (max - min + 1)) + min;
    },
    async bubbleSort() {
      let len = this.array.length;
      let checked;
      do {
        checked = false;
        for (let i = 0; i < len; i++) {
          if (this.array[i] > this.array[i + 1]) {
            let tmp = this.array[i];
            this.array[i] = this.array[i + 1];
            this.array[i + 1] = tmp;
            // sleep - to visualize / see the changes
            await this.sleep();
            checked = true;
          }
        }
      } while (checked);
    },
    sleep() {
      return new Promise((resolve) => setTimeout(resolve, 100));
    },
  },
  created() {
    this.fillArray();
  }
});

app.mount("#app");
