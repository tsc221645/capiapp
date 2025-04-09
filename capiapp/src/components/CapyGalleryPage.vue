<template>
    <div class="capy-gallery-page">
        <header class="header">
            <div class="logo" @click="toggleMenu">
                <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
                <span>CaypGallery</span>
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

    </div>

</template>


<script>
import MenuTab from './MenuTab.vue'
export default{
    components: { MenuTab },
    name: 'CapyGalleryPage',
    data() {
        return {
        darkMode: false,
        menuVisible: false,
        menuItems: [
          { label: 'main page', route: '/' },
          { label: 'capybara facts', route: '/capyfacts' },
          { label: 'galery', route: '/capygallery' },
          { label: 'capygames', route: '/capygames' },
          { label: 'capyquotes', route: '/capyquotes' },
          { label: 'meditating with capy', route: '/meditate' }
        ]
        };
    },
    methods: {
        toggleTheme() {
        const theme = this.darkMode ? 'dark' : 'light';
        document.documentElement.setAttribute('data-theme', theme);
        },
        goToLogin() {
            this.$router.push({ name: 'LoginPage' });
        },
        toggleMenu() {
          this.menuVisible = !this.menuVisible;
        }
    },
    mounted() {
        this.toggleTheme();
    }
}

</script>

<style>
.capy-gallery-page {
    background-color: var(--bg-color);
    height: 100dvh;
    font-family: 'Libre Baskerville', serif;
}
</style>