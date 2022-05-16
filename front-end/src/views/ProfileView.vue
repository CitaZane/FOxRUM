<template>
    <SideProfile v-if="isLoggedIn" :resp="resp"/>
    <MainProfile v-if="isLoggedIn" :resp="resp"/>
</template>

<script>
// @ is an alias to /src
import SideProfile from '@/components/sidebar/SideProfile.vue'
import MainProfile from '@/components/main/MainProfile.vue'


import {mapState , mapActions} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
  name: 'ProfileView',
  components: {
    SideProfile, 
    MainProfile
  },
  props:["resp"],
  computed:{
    ...mapState(useMainStore,["ws"]),
    ...mapState(useMainStore,["isLoggedIn"]),
    ...mapState(useMainStore,["name"])
  },
  methods:{
    ...mapActions(useMainStore, {resetStore : 'reset'}),
    fetchUserPosts(){
    this.ws.send(JSON.stringify({
          action: 'get-user-posts',
          message: "0" // offset
      }))
    }
  },
  created(){
    if(typeof this.isLoggedIn == 'undefined' ||!this.isLoggedIn){
      this.resetStore()
      this.$router.push({name:'login'})
      return
    }
    this.fetchUserPosts()
  }
}
</script>

<style scoped>

</style>