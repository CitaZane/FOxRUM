<template>
       <div class="side-wrapper">
        <div class="side-container">
            <div class="side-content">
            <router-link :class="[path =='login'? 'auth-link-active': 'auth-link']" to="/login">
                <p :class="[path =='login'? 'auth-title-active': 'auth-title']">Log in</p>
            </router-link>

            <router-link :class="[path=='signup'? 'auth-link-active': 'auth-link']" to="/signup">
                <p :class="[path =='signup'? 'auth-title-active': 'auth-title']">Sign up</p>
            </router-link>

            
            </div>

        </div>
    </div>
</template>

<script>
export default {
    props:["path"],

}
</script>


<style scoped>
.side-container{
    height: 90vh;
    margin-top: 5vh;
    padding: 0vh 2vw 0vh 2vw;
    overflow: hidden;
    position: relative;
}
.side-content{
    position: absolute;
    width: 14vw;
    top: 50%;
    transform: translateY(-50%);
}
@media (max-width: 900px){
    .side-content{
        width: 80%;
    }
}

.auth-link{
    background-color: var(--content-color);
    width: 100%;
    display:flex;
    justify-content: space-between;
    
    border-radius: 0.5rem;
    text-decoration: none;
    padding: 0.5vh 3vh 0.5vh 3vh;
    margin-bottom: 5px;
}
.auth-link-active{
    background-color: var(--main-color);
    width: 100%;
    /* height: 28px; */
    display:flex;
    justify-content: space-between;
    border-radius: 0.5rem;
    text-decoration: none;
    padding: 0.5vh 3vh 0.5vh 3vh;
    margin-bottom: 5px;
}
.auth-title{
   color: var(--main-color);
    font-size: var(--font-m); 
}
.auth-title-active{
   color: var(--accent-text-color);
    font-size: var(--font-m);
    font-weight: bold; 
}
.auth-link:hover{
    background-color: var(--content-color-active);
}

</style>