<template>
    <router-view />
</template>

<script>


export default {
  name: 'App',

  
}
</script>

<style>
@import "@/assets/themes.css";

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
