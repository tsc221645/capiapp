<template>
    <div class="profile-page">
        <div class="home-page">
            <header class="header">
                <div class="logo" @click="toggleMenu">
                    <img src="../assets/capybara_logo.png" alt="Capybara Logo" class="logo-img" />
                    <span>Profile Page</span>
                </div>

                <div class="header-controls">
                    <p>logged in capy</p>

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

            <div class="elements-container">
                <div class="image-container">


                </div>
                <div class="main-container">
                    <p></p>

                </div>
            </div>

            

        </div>
    </div>


</template>


<script>
import MenuTab from './MenuTab.vue'
export default{
    components: { MenuTab },
    name: 'ProfilePage',
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
.profile-page {
    background-color: var(--bg-color);
    height: 100dvh;
    font-family: 'Libre Baskerville', serif;
}

.elements-container{
    display:flex;
}

.image-container{
    background-color: var(--primary-color);
    height: 300px;
    width: 300px;
    border-radius: 50%;
    transform: translateX(10vw);
}


/*container for inputs*/
.main-container{
    background-color: var(--primary-color);
    height: 70vh;
    width: 58vw;
    transform: translateX(20vw);
    padding-left: 20px;
    padding-top:20px;
    border-radius: 20px;
    box-shadow: 0px 4px 8px var(--text-color);
}

</style>