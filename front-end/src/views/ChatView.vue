<template>
        <SideChat v-if="isLoggedIn" :resp="resp" :chatRoom="chatRoom"/>
        <MainChat v-if="isLoggedIn" :resp="resp" :chatRoom="chatRoom"/>
</template>
<script>
// @ is an alias to /src
import SideChat from '@/components/sidebar/SideChat.vue'
import MainChat from '@/components/main/MainChat.vue'

import {mapState, mapActions} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
  name: 'ChatView',
  components: {
    SideChat,
    MainChat
  },
  props:["resp"],
  methods:{
    ...mapActions(useMainStore, {resetStore : 'reset'}),
    getAllUsers() {
      this.ws.send(JSON.stringify({
        action: 'get-all-users',
      }))
    }
  },
  computed:{
    chatRoom(){
        return this.$route.params.id
    },
    ...mapState(useMainStore,["ws"]),
    ...mapState(useMainStore,["isLoggedIn"]),
  },
  created(){
    if(typeof this.isLoggedIn == 'undefined' ||!this.isLoggedIn){
      this.resetStore()
      this.$router.push({name:'login'})
      return
    }
    this.getAllUsers()
  }
}
</script>