<template>
    <main class="main-wrapper">
        <div class="main-container">
            <div class="main-header">
            <h2 class="main-title">It looks like you are lost...</h2>
            <hr class="divider">
            </div>

            <div :class="drawingClass" class="main-body">
                <div :class="errorClass"></div>
            </div>

        </div>
    </main>
</template>

<script>
import { mapState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'

    export default {
        computed:{
        ...mapState(useMainStore,["darkMode"]),
    drawingClass(){
      return (this.darkMode)? "drawing-dark": "drawing"
        },
        errorClass(){
      return (this.darkMode)? "error-dark": "error"
        }
     
     }
    }
</script>

<style scoped>
.drawing{
   background-image: url("@/assets/404-forest.png");
}
.drawing-dark{
   background-image: url("@/assets/404-forest_dark.png");
}
.drawing, .drawing-dark{
    height: 100%;
    background-repeat: no-repeat;
    background-size: cover;
    display: flex;
    justify-content: center;
    align-items: center;

}
.error{
   background-image: url("@/assets/404.png");
}
.error-dark{
   background-image: url("@/assets/404_dark.png");
}
.error, .error-dark{
    height: 65%;
    aspect-ratio: 2/1;
    background-repeat: no-repeat;
    background-size: contain;
      align-self: center;


}
</style>