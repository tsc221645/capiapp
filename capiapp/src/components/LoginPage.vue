<template>
    <div class = "login-page">
        <div class="login-wrapper">
            <div class="login-card">
            <h2>Login</h2>
            <input type="text" v-model="username" placeholder="username" />
            <input type="password" v-model="password" placeholder="password" />
            <button @click="handleLogin">Login</button>
            
            <div  class="logo signup-text">
                <span><router-link to="/signup">Don't have an account? Sign Up</router-link></span>
                <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
            </div>
        </div>
    </div>
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

html, body {
  margin: 0;
  padding: 0;
  background-color: var(--bg-color);
  height: 100%;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}


.login-page {
  background-color: var(--bg-color);
  height: 100dvh;
  font-family: 'Libre Baskerville', serif;
}

.login-wrapper {
  
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-card {
  background: var(--primary-color);
  align-items: center;
  padding: 40px 30px;
  border-radius: 20px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.3);
  width: 40vw;
  height: 70vh;
  justify-content: center;
  text-align: center;
}

.login-card h2 {
  margin-top: 20px;
  margin-bottom: 70px;
  color: var(--text-color);
  letter-spacing: 0.1em;
  font-size: clamp(2.5rem, 80vw, 1rem);
}

.login-card input {
  align-items: center;
  width: 80%;
  height: 10%;
  font-family: 'Libre Baskerville', serif;
  font-size: 18px;
  padding: 15px;
  margin-bottom: 30px;
  border: none;
  border-radius: 8px;
  box-shadow: inset 0 1px 3px rgba(0,0,0,0.1);
}

.login-card button {
  margin-top: 50px;;
  background: var(--accent-color-1);
  color: var(--text-color);
  border: none;
  padding: 10px;
  border-radius: 15px;
  width: 40%;
  height: 10%;
  cursor: pointer;
  box-shadow: 0px 4px 8px var(--text-color);
  font-size: 18px;
  font-family: 'Libre Baskerville', serif;
  transition: background 0.2s;
}

.login-card button:hover {
  background: var(--accent-color-4);
    color: var(--bg-color);
}

.signup-text {
  margin-top: 15px;
  font-size: 14px;
  color: var(--text-color);
  display: flex;
  justify-content: center; 
  align-items: center;     
  gap: 8px;                
}
.signup-text a{
  color: var(--accent-color-4);
  font-weight: bold;
  text-decoration: none;
}

.signup-text a:hover {
  text-decoration: underline;
}


</style>