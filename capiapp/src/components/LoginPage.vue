<template>
    <div class="login-container">
        <form @submit.prevent="handleLogin">
        <h2>Login</h2>
        <input type="text" v-model="username" placeholder="username" required />
        <input type="password" v-model="password" placeholder="password" required />
        <button type="submit">Login</button>
        <p>
            Don’t have an account? 
            <router-link to="/signup">Sign Up</router-link>
        </p>
        </form>
  </div>

</template>

<script>
export default{
    name: 'LoginPage',
    data(){
        return{
            username:'',
            email:'',
            password:''
        }
    },
    methods:{
        goToSignUp() {
            this.$router.push({ name: 'SignUp' });
        }
    }
    

}

</script>

<style>
.login-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background-color: #f0f0f0;
}


</style>