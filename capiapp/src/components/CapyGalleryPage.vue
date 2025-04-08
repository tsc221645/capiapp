<template>
    <div class="capy-gallery-page">
        <header class="header">
            <div class="logo">
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
        </header>

    </div>

</template>


<script>
export default{
    name:'CapyGalleryPage',
    methods:{
        goToLogin(){
            this.$router.push({name:'LoginPage'})
        },
        toggleTheme() {
            const theme = this.darkMode ? 'dark' : 'light';
            document.documentElement.setAttribute('data-theme', theme);
        }
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