<template>
    <head>

    </head>
    <body>
        <p>Welcome to the Home Page</p>
    </body>


</template>


<script>
export default{
    name: 'HomePage',
    mounted(){
        let user = localStorage.getItem('user-info'); //check if user is already logged in
        if(user){
            this.$router.push({name:'SignUp'});//redirect to the signup page if user is already logged in
        }
    }
}


</script>


<style>



</style>