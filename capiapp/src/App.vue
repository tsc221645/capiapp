<template>

  <head>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&display=swap" rel="stylesheet">
  </head>
    <router-view />
</template>

<script>


export default {
  name: 'App',

  
}
</script>

<style>
@import "@/assets/themes.css";
@import url('https://fonts.googleapis.com/css2?family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&display=swap');

.app {
  background-color: var(--bg-color);
  color: var(--text-color);
  min-height: 100vh;
  padding: 1rem;
}

button {
  background-color: var(--accent-color-1);
  border: none;
  color: var(--text-color);
  padding: 0.5rem 1rem;
  cursor: pointer;
}


</style>
