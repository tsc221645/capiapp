import HomePage from "./components/HomePage.vue";
import SignUpPage from "./components/SignUpPage.vue";
import LoginPage from "./components/LoginPage.vue";
import CapyFactsPage from "./components/CapyFactsPage.vue";
import CapyGalleryPage  from "./components/CapyGalleryPage.vue";
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
    {
      name: "CapyFactsPage",
      component: CapyFactsPage,
      path: "/capyfacts",
    },
    {
      name:"CapyGalleryPage",
      component: CapyGalleryPage,
      path:"/capygallery",
    },
  ];
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;