import HomePage from "./components/HomePage.vue";
import SignUpPage from "./components/SignUpPage.vue";
import LoginPage from "./components/LoginPage.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
      name: "HomePage",
      component: HomePage,
      path: "/",
    },
    {
      name: "SignUpPage",
      component: SignUpPage,
      path: "/signup",
    },
    {
        name: "LoginPage",
        component: LoginPage,
        path: "/login",
      },
  ];
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;