<template>
    <div class="profile-page">
        <div class="home-page">
            <header class="header">
                <div class="logo">
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
            </header>

        </div>
    </div>


</template>


<script>
export default {
    name: "ProfilePage",
}

</script>

<style>
.profile-page {
    background-color: var(--bg-color);
    height: 100dvh;
    font-family: 'Libre Baskerville', serif;
}
</style>