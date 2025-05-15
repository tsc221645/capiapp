<template>
    <div class="capy-facts-page">
        <header class="header">
            <div class="logo" @click="toggleMenu">
                <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
                <span>CapyFacts</span>
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
            <transition name="fade-menu">
                <MenuTab :show="menuVisible" :items="menuItems" />
            </transition>
        </header>
        <div class="facts-container">
            <div v-if="capy" class="fact-card">
            <div class="card-header">
                <img src="https://upload.wikimedia.org/wikipedia/commons/6/6b/Capybara_%28Hydrochoerus_hydrochaeris%29.JPG" alt="Capybara" class="capy-img" />
                <h2>{{ capy.name }}</h2>
            </div>
            <div class="card-body">
                <p><strong>Nombre científico:</strong> {{ capy.taxonomy?.scientific_name }}</p>
                <p><strong>Ubicación:</strong> {{ capy.locations?.join(', ') }}</p>
                <p><strong>Dato curioso:</strong> {{ capy.characteristics?.interesting_fact }}</p>
            </div>
            </div>



        </div>

    </div>

</template>


<script>
import MenuTab from './MenuTab.vue'
import axios from 'axios'

export default {
  components: { MenuTab },
  name: 'CapyFactsPage',
  data() {
    return {
      darkMode: false,
      menuVisible: false,
      capy: null,
      menuItems: [
        { label: 'main page', route: '/' },
        { label: 'capybara facts', route: '/capyfacts' },
        { label: 'galery', route: '/capygallery' },
        { label: 'capygames', route: '/capygames' },
        { label: 'capyquotes', route: '/capyquotes' },
        { label: 'meditating with capy', route: '/meditate' }
      ]
    }
  },
  methods: {
    toggleTheme() {
      const theme = this.darkMode ? 'dark' : 'light'
      document.documentElement.setAttribute('data-theme', theme)
    },
    goToLogin() {
      this.$router.push({ name: 'LoginPage' })
    },
    toggleMenu() {
      this.menuVisible = !this.menuVisible
    },
    async fetchCapyFact() {
        const res = await axios.get("http://backend:3000/capyfacts");
        this.capy = res.data[0];
    }

  },
  mounted() {
    this.toggleTheme()
    this.fetchCapyFact()
  }
}
</script>



<style>
.capy-facts-page {
    background-color: var(--bg-color);
    height: 100dvh;
    font-family: 'Libre Baskerville', serif;
}


.facts-container{
    background-color: var(--bg-color);
    display:block;
    font-family: 'Libre Baskerville', serif;
}


.fact-card {
  background-color: var(--card-bg, #fff);
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.1);
  padding: 20px;
  margin: 30px auto;
  max-width: 600px;
  color: var(--text-color, #333);
  transition: transform 0.3s ease;
}
.fact-card:hover {
  transform: scale(1.01);
}
.card-header {
  display: flex;
  align-items: center;
  gap: 16px;
}
.capy-img {
  width: 100px;
  height: 100px;
  border-radius: 12px;
  object-fit: cover;
}
.card-body {
  margin-top: 16px;
}
.card-body p {
  margin: 6px 0;
}

</style>