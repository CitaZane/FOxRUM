import { defineStore } from 'pinia'

export const useMainStore = defineStore({
    id: "main",
  state:()=>({
          isLoggedIn:false,
          ws: null,
          name: "",
          chatNotification: false,
          darkMode: false,
  }),
  actions: {
    reset() {
      this.isLoggedIn =false,
      this.ws= null,
      this.name= "",
      this.chatNotification= false
    }
  },
  persist: true, // plugin - presist through refresh
  beforeRestore: context => {
      console.log('Before hydration...', context)
    },
    afterRestore: context => {
      console.log('After hydration...', context)
    },
});