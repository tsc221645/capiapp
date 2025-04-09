<template>
    <div class="menu-tab" v-if="show">
      <div class="menu-content">
        <h3>menu</h3>
        <ul>
          <li v-for="(item, index) in items" :key="index" @click="navigate(item)">
            {{ item.label }}
            <span class="arrow">›</span>
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'MenuTab',
    props: {
      show: {
        type: Boolean,
        required: true
      },
      items: {
        type: Array,
        required: true
      }
    },
    methods: {
      navigate(item) {
        if (item.route) {
          this.$router.push(item.route);
        } else if (item.action) {
          item.action(); // permite ejecutar funciones si no es ruta
        }
      }
    }
  }
  </script>
  
  <style scoped>
  .menu-tab {
    position: absolute;
    top: 80px;
    left: 10px;
    background: var(--primary-color);
    border-radius: 20px;
    padding: 20px;
    height:90vh;
    width: 35vw;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    z-index: 1000;
  }
  
  .menu-content h3 {
    margin: 0 0 10px 0;
    font-weight: bold;
  }
  
  .menu-content ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  
  .menu-content li {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 0;
    cursor: pointer;
    transition: background 0.2s;
  }
  
  .menu-content li:hover {
    background-color: #e8d9ca;
    border-radius: 5px;
  }
  
  .arrow {
    font-weight: bold;
  }

  
  </style>
  