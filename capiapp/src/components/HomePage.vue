<template>
    <div class="home-page">
    <header class="header">
      <div class="logo">
        <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
        <span>Capy App</span>
      </div>

      <div class="header-controls">
        <button class="login-button" @click="goToLogin">Login</button>

        <div class="theme-toggle">
          <label class="switch">
            <input type="checkbox" v-model="darkMode" @change="toggleTheme" />
            <span class="slider"></span>
          </label>
          
        </div>
      </div>
    </header>
    
      <main class="main-content">
        <div class="left">
          <h1>
            A site<br />
            dedicated<br />
            only to<br />
            capybaras
          </h1>
        </div>

        <div class="right actions ">
          <button class="action-btn green">Chat with other capybara <br/>enthusiasts</button>
          <button class="action-btn brown">Discover new facts</button>
          <button class="action-btn beige">Create your own <br/>capi-community</button>
        </div>
      </main>
    </div>
</template>


<script>
export default{
    name: 'HomePage',
    data() {
        return {
        darkMode: false
        };
    },
    methods: {
        toggleTheme() {
        const theme = this.darkMode ? 'dark' : 'light';
        document.documentElement.setAttribute('data-theme', theme);
        },
        goToLogin() {
            this.$router.push({ name: 'LoginPage' });
        }
    },
    mounted() {
        this.toggleTheme();
    }
}


</script>


<style >

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


.home-page {
  background-color: var(--bg-color);
  height: 100dvh;
}


.header {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: var(--primary-color);
  padding: 0.8rem 2rem; 
  border-radius: 12px;
  margin-bottom: 2rem;
  box-sizing: border-box;
  box-shadow: 0px 4px 8px var(--text-color);
  font-family: 'Libre Baskerville', serif;
}


.header-controls {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.login-button {
  background-color: #75502D;
  color: white;
  padding: 0.4rem 1rem;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

/* Dark mode switch */
.theme-toggle {
  display: flex;
  align-items: center;
  gap: 0.4rem;
}

.switch {
  position: relative;
  display: inline-block;
  width: 42px;
  height: 24px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: var(--toggle-background-color);
  border-radius: 24px;
  transition: 0.3s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  border-radius: 50%;
  transition: 0.3s;
}

input:checked + .slider {
  background-color: var(--toggle-background-color);
}

input:checked + .slider:before {
  transform: translateX(18px);
}

.theme-icon {
  font-size: 1.2rem;
  transform: translateX(18px);
}


.logo {
  display: flex;
  align-items: center;
  font-size: 1.5rem;
}

.logo-img {
  width: 40px;
  height: 40px;
  margin-right: 10px;
}

.actions{
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 100%;
    
}

.action-btn {
  font-size: 2rem;
  font-family: 'Libre Baskerville', serif;
  padding: 1rem;
  border: none;
  border-radius: 10px;
  text-align: left;
  cursor: pointer;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  color: white;
  height: 20vh;
  width: 40vw;
}

.green {
  background-color: var(--accent-color-3);
}

.brown {
  background-color: var(--accent-color-4);
}

.beige {
  background-color: var(--primary-color);
}


.main-content {
  display: flex;
  background-color: var(--bg-color);
  font-family: 'Libre Baskerville', serif;
  letter-spacing: 0.1em;
  justify-content: space-center;
  align-items: flex-start;
  font-size: clamp(3rem, 6vw, 4rem);
  padding: 2rem;
  gap: 4rem;
  max-width: 1500px;
  margin: 0 auto;
  color: var(--text-color);
  text-shadow: 1px 1px 4px var(--text-color);
}


.left {
  flex: 1;
  
  margin-right: auto;
}

.left h1 {
  font-weight: 100;
  line-height: 1.1;
}

.right {
  
  flex: 1;
  
}

</style>