import HomePage from "./components/HomePage.vue";
import SignUp from "./components/SignUp.vue";
import LogIn from "./components/Login.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
      name: "HomePage",
      component: HomePage,
      path: "/",
    },
    {
      name: "SignUp",
      component: SignUp,
      path: "/signup",
    },
    {
        name: "LogIn",
        component: LogIn,
        path: "/login",
      },
  ];
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;