<template>
  <div class="single-comment">
      <div class="single-comment-stats">
          <p>{{ commentAuthor() }}</p>
          <p>Barked back at: {{ commentTime() }}</p>
      </div>
    <p class="single-comment-content">{{ comment.content }}</p>
  </div>
</template> 

<script>
export default {
    props:["comment"],
    methods:{
       commentAuthor() {
      return "Fox: "+ this.comment.author;
    },
    commentTime(){
        const date = new Date(Number(this.comment.time)*1000)
        const today = new Date()
        let minutes= (Number(date.getMinutes())<10)? "0"+date.getMinutes(): date.getMinutes();
        return (today.getDate() === date.getDate())? date.getHours()+":"+minutes: date.getDate()+"/"+date.getMonth()+"/" + date.getFullYear()    
    } 
    }
}
</script>

<style scoped>
    .single-comment{
        background-color: var(--content-color);
        border-radius: 0.75rem;
        margin-bottom: 1.5vh;
        padding: 0.2rem 2rem 0.2rem 2rem;
    }
    .single-comment-stats{
        display:flex;
        justify-content: space-between;
    }
    .single-comment-stats p{
        font-size: var(--font-xs);
    }
    .single-comment-content{
        text-align: left;
        margin-top: 0.3rem;
        font-size: var(--font-s);
        font-weight: 600;
    }
</style>