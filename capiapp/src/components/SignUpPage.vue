<template>
    <div class = "signup-page">
        <div class="signup-wrapper">
            <div class="signup-card">
            <h2>Sign Up</h2>
            <input type="text" v-model="username" placeholder="username" />
            <input type="text" v-model="email" placeholder="email" />
            <input type="password" v-model="password" placeholder="password" />
            <button @click="signup">Sign Up</button>
            
            <div  class="logo signup-text">
                <span><router-link to="/login">Already have an account? Login</router-link></span>
                <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
            </div>
        </div>
    </div>
    </div>

</template>


<script>
export default{
    name: 'SignUpPage',
    data(){
        return{
            username:'',
            email:'',
            password:''
        }
    },
    methods:{ 
      async signup() {
        try {
          const res = await fetch('http://localhost:3000/signup', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              username: this.username,
              email: this.email,
              password: this.password
            })
          })

          if (!res.ok) {
            const errorData = await res.json()
            throw new Error(errorData.error || 'Error al registrar')
          }

          alert('Usuario creado correctamente. Inicia sesión.')
          this.$router.push({ name: 'Login' })
        } catch (err) {
          this.error = err.message
        }
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


.signup-page {
  background-color: var(--bg-color);
  height: 100dvh;
  font-family: 'Libre Baskerville', serif;
}

.signup-wrapper {
  
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.signup-card {
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

.signup-card h2 {
  margin-top: 20px;
  margin-bottom: 70px;
  color: var(--text-color);
  letter-spacing: 0.1em;
  font-size: clamp(2.5rem, 80vw, 1rem);
}

.signup-card input {
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

.signup-card button {
  margin-top: 30px;;
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

.signup-card button:hover {
  background: var(--accent-color-4);
    color: var(--bg-color);
}

.signup-text {
  margin-top: 25px;
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