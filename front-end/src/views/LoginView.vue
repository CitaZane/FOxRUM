<template>
    <SideLogin :path="path"/>
    <MainLogin :resp="resp"/>
</template>

<script>
// @ is an alias to /src
import SideLogin from '@/components/sidebar/SideLogin.vue'
import MainLogin from '@/components/main/MainLogin.vue'

import {mapState} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
  name: 'LoginView',
  components: {
    SideLogin, 
    MainLogin
  },
  computed:{
    ...mapState(useMainStore,["isLoggedIn"]),
  },
  props:["resp"],
  data: function(){
    return{
      path: "login"
    }
  },
  created(){
     if(this.isLoggedIn){
      this.$router.push({name:'feed'})
    }
  }
}
</script>
