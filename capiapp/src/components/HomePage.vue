<template>
    <header class="header">
        <div class="logo">
            <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
            <span>Capi App</span>
        </div>
    </header>
    <body>
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
        

        
    </body>
    <footer>

    </footer>


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
        }
    },
    mounted() {
        this.toggleTheme();
    }
}


</script>


<style>

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

.login-button{
    background-color: var(--login-btn);
    color: var(--text-color);
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