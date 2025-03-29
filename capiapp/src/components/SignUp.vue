<template>
    <head>
        <h1>Sign Up</h1>
    </head>
    <body>
    </body>

</template>


<script>
export default{
    name: 'SignUp',
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


<style>



</style>