const app = Vue.createApp({
    data(){
        return {
            message: "Welcome!"
        };
    },
    methods: {
        sayHello(){
            this.message = "Hello!"
        }
    }
});

app.mount("#app");