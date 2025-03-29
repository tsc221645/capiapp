<template>
    <head>
        <h1> Login </h1>
    </head>

</template>

<script>
export default{
    name: 'LogIn',
    data(){
        return{
            username:'',
            email:'',
            password:''
        }
    },
    mounted(){
        let user = localStorage.getItem('user-info'); //check if user is already logged in
        if(user){
            this.$router.push({name:'HomePage'});//redirect to the home page if user is already logged in            }
        }
    }

}

</script>