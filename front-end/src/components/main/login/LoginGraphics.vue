<template>
<div class="login-graphics">
    <div id="foxrum"></div>
    <div :class="foxDrawing" id="fox-drawing"></div>
    <p class="login-quote-one">ALL IN ALL, ISN'T IT ALL</p>
    <p class="login-quote-one">ABOUT FOXES AND RUM</p>
    <p class="login-quote-two">Just pour some and then lets talk</p>
</div>
   
</template>


<script>
import { mapState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'
export default {
    computed:{
    ...mapState(useMainStore,["darkMode"]),
    foxDrawing(){
        return (this.darkMode)? "fox-drawing-dark": "fox-drawing"
    }
  },
}
</script>
<style scoped>
.login-graphics{
    height: 100%;
    width: 36vw;
    flex-grow: 1;
    align-items: center;
    display:flex;
    justify-content: space-between;
    flex-direction: column;
    border-left: 1px solid var(--content-color-dark);
}
@media (max-width: 900px){
    .login-graphics{
        display: none;
    }
}
#foxrum{
    flex-grow: 2;
    background-image: url("@/assets/foxrum.svg");
    background-repeat: no-repeat;
      background-size: contain;
      height: 20vh;
      aspect-ratio: 2/1;
      max-width: 32vw;
}
.fox-drawing{
    flex-grow: 4;
    background-image: url("@/assets/fox-welcome.png");
    background-repeat: no-repeat;
      background-size: contain;
      height: 50vh;
      aspect-ratio: 1/1.2;
      max-width: 32vw;
}
.fox-drawing-dark{
    flex-grow: 4;
    background-image: url("@/assets/fox-welcome_dark.png");
    background-repeat: no-repeat;
      background-size: contain;
      height: 50vh;
      aspect-ratio: 1/1.2;
      max-width: 32vw;
}
.login-quote-one{
    color:var(--accent-color-light);
    font-weight: bold;
    font-size: var(--font-m);
    flex-grow: 1;

}
.login-quote-two{
    color:var(--accent-color-light);
    font-weight: lighter;
    font-style: italic;
    font-size: var(--font-m);
    flex-grow: 1;

}
</style>