<template>
  <router-link :to="postURL()" class="single-post" >
      <div class="single-post-stats">
          <p>{{ postUser() }}</p>
          <p>Howled at: {{ postTime() }}</p>
      </div>
    <p class="single-post-content">{{ post.title }}</p>
  </router-link>
</template> 

<script>
export default {
  props: ["post"],
  methods: {
    postUser() {
      return "Fox: "+ this.post.author;
    },
    postTime(){
      const date = new Date(Number(this.post.time)*1000)
        const today = new Date()
        let minutes= (Number(date.getMinutes())<10)? "0"+date.getMinutes(): date.getMinutes();
        return (today.getDate() === date.getDate())? date.getHours()+":"+minutes: date.getDate()+"/"+date.getMonth()+"/" + date.getFullYear()   
    },
    postURL(){
      return "/post/"+this.post.id;
    }
  }};
</script>
<style scoped>
    .single-post{
        background-color: var(--content-color);
        border-radius: 0.75rem;
        margin-bottom: 1.5vh;
        padding: 0.2rem 2rem 0.2rem 2rem;
        display:block;
        text-decoration: none;
        color: var(--main-color)

    }
    .single-post-stats{
        display:flex;
        justify-content: space-between;
    }
    .single-post-stats p{
        font-size: var(--font-xs);
    }
    .single-post-content{
        text-align: left;
        margin-top: 0.3rem;
        font-size: var(--font-m);
        font-weight: 600;
    }
    .single-post:hover{
      background-color: var(--content-color-active);
    }
</style>