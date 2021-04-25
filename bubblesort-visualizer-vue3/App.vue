<template>
  <h1>Sorting Algorithm Visualizer</h1>
  <button @click="bubbleSort()">Bubble Sort</button>
  <button @click="fillArray()">Shuffle</button>

  <div class="container">
    <!-- Loop through array numbers and display their height. -->
    <div class="bar" v-for="(number,index) in array" :key="index" :style="{height: number + 'px'}">
    </div>
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      array: [],
      size: 30,
    }

  },
  components: {
  },
  methods:{
    fillArray(){
      // populate array with random numbers between 5 and 750
      this.array = [];
      for (let i =0; i< this.size; i++) {
        this.array.push(this.getRndInteger(5, 750))
      }
    },
    getRndInteger(min, max) {
      return Math.floor(Math.random() * (max - min + 1) ) + min;
    },
    async bubbleSort(){
      let len = this.array.length
      let checked;
      do {
        checked = false;
        for(let i =0;i<len;i++){
          if(this.array[i] > this.array[i+1]){
            let tmp = this.array[i];
            this.array[i] = this.array[i+1];
            this.array[i+1] = tmp;
            // sleep - to visualize / see the changes
            await this.sleep()
            checked = true
          }
        }
      } while(checked)
    },
    sleep() {
      return new Promise((resolve) => setTimeout(resolve, 500));
    },
  },
  created() {
    this.fillArray();
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.container {
  position: fixed;
  left: 100px;
}

.bar {
  width: 20px;
  background-color: black;
  display: inline-block;
  margin: 0 5px;
}
</style>
