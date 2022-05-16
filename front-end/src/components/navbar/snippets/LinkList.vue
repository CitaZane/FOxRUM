<template>
    <div class="nav-link-container">
        <nav-link v-for="link in links" :key="link.id" :link="link"></nav-link>
        <a href="#" class="nav-link" v-if="isLoggedIn" @mouseup="logout">
            <span class="nav-icon icon-logout" alt="link" id="nav-link-logout"></span>
        </a>
    </div>
</template>

<script>
    import NavLink from "@/components/navbar/snippets/NavLink.vue";
    import { mapWritableState, mapActions } from 'pinia' //helper for store

    import {useMainStore} from '@/stores/main.js'

    export default {
    components: { NavLink },
    props: ["links"],
    computed:{
        ...mapWritableState(useMainStore,['isLoggedIn']),
        ...mapWritableState(useMainStore,['name']),
        ...mapWritableState(useMainStore,['ws']),
        ...mapWritableState(useMainStore,['chatNotification'])
  },
    methods:{
        ...mapActions(useMainStore, {resetStore : 'reset'}),
        logout(){
/* ------------------------- destroy session cookie ------------------------- */
                document.cookie = "session-id=;max-age=-1"
/* ------------------------- stop ws conn for client ------------------------ */
            this.ws.send(JSON.stringify({
            action: 'user-logout',
            message: this.name
            }))
            // this.ws.close() //close connection
/* ------------------------------- reset store ------------------------------ */
            this.resetStore()
            // localStorage.removeItem("main")
/* ---------------------------- redirect to login --------------------------- */
            this.goToLogin();
        },
        goToLogin(){
            this.$router.push({name:'login'})
        }
    }
    };
</script>

<style scoped>
.nav-link-container{
        position: absolute;
        width:4.5vw;
        top: 50%;
        -ms-transform: translateY(-50%);
        transform: translateY(-50%);
    }

</style>