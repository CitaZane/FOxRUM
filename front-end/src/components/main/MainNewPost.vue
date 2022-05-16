<template>
    <main class="main-wrapper">
        <div class="main-container">
        <h2 class="main-title">What are you howling about?</h2>
        <hr class="divider">
        <form id="new-post-form">
            <div class="form-wide">
                <!-- /* ---------------------------------- title --------------------------------- */ -->
                 <div :class="v$.newPost.title.$error? errorClass : activeClass">
                    <label for="form-post-title">Howl title</label>
                    <input v-model="v$.newPost.title.$model" id="form-post-title" placeholder="Title..." class="input-field">
                    <span v-if="v$.newPost.title.$error" class="err-msg">
                            {{ v$.newPost.title.$errors[0].$message }}
                    </span>
                </div>
                
                <!-- /* --------------------------------- content -------------------------------- */ -->
                <div :class="v$.newPost.content.$error? errorClass : activeClass">
                    <label for="form-post-content">Scream it out</label>
                    <textarea v-model="v$.newPost.content.$model" id="form-post-content" placeholder="Description..." class="input-field"></textarea>
                    <span v-if="v$.newPost.content.$error" class="err-msg">
                            {{ v$.newPost.content.$errors[0].$message }}
                    </span>
                </div>   
            </div>
            <div :class="formNarrowClass">

                <div id="form-narrow-one">
                    <!-- /* -------------------------------- category -------------------------------- */ -->
                    <div :class="v$.newPost.category.$error? errorClass : activeClass">
                        <label for="form-post-category">Howl category</label>
                        <input v-model="v$.newPost.category.$model" id="form-post-category" list="categoryList" placeholder="Category..." class="input-field" @keyup.enter.exact="post()">
                        <span v-if="v$.newPost.category.$error" class="err-msg">
                                {{ v$.newPost.category.$errors[0].$message }}
                        </span>
                        <datalist id="categoryList">
                        <option v-for="category in categoryList" :key="category.id" :value="category.name"></option>
                        </datalist>
                    </div>
                    <div @mousedown="post()" id="post-new-post-btn" class="btn">Howl</div>
                </div>
                <div :class="foxDrawing" id="form-narrow-two"></div>

            </div>

            

        </form>
        </div>
    </main>
</template>

<script>
import { mapState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'
/* ------------------------------- validation ------------------------------- */
import useValidate from "@vuelidate/core";
import { required, maxLength, helpers} from "@vuelidate/validators";

export default {
    props:["resp"],
    data: function(){
        return{
            v$: useValidate(),
            newPost:{
                title:"",
                content:"",
                category:"",
                author:""
            }, 
            activeClass: "input-group",
            errorClass: "input-group-error"
        }
    },
    computed:{
        ...mapState(useMainStore,["ws"]),
        ...mapState(useMainStore,["name"]),
        ...mapState(useMainStore,["darkMode"]),
        foxDrawing(){
            return (this.darkMode)? "fox-drawing-dark": "fox-drawing"
        },
        formNarrowClass(){
            return (this.darkMode)? "form-narrow-dark": "form-narrow"
        },
        categoryList(){
            return this.prepereCategoryList()
        }
    },
    methods:{
        prepereCategoryList(){
             let categories = []
            let id = 1
            this.resp.categories.forEach(cat => {
                let category = {name: cat, id: id}
                categories.push(category)
                id++
            });
            return categories
        },
        async post(){
            const isFormCorrect = await this.v$.$validate()
            if (!isFormCorrect) return
            this.newPost.author = this.name
            this.ws.send(JSON.stringify({
                action: 'new-post',
                message: JSON.stringify(this.newPost)
            })),
            this.goToFeed(this.newPost.category)
        },
        goToFeed(category) {
            this.$router.push({name:'category', params:{name:`${category}`}})
        },
    },
    validations(){
        return{
            newPost:{
                title:{
                    required: helpers.withMessage("Post without a title? haven't heard of that", required),
                    maxLength: helpers.withMessage("This title is a bit too loooooong!", maxLength(255))
                },
                content:{
                    required: helpers.withMessage("Content is required. Content is everything!", required),
                },
                category:{
                    required: helpers.withMessage("Please choose or create a category!", required),
                    maxLength: helpers.withMessage("This category is a bit too loooooong!", maxLength(60))
                }
            }
        }
    }
}
</script>


<style scoped>
.main-title{
    font-size: var(--font-l);
    font-weight: bold;
    color: var(--main-color);
}
#new-post-form{
    text-align: left;
    font-size:var(--font-m);
    font-weight: 500;
    height: calc(90vh - 10vh - var(--font-l))
}

.form-narrow, .form-narrow-dark{
    position: relative;
    display:flex;
    justify-content: space-between;
    overflow: hidden;
}
.form-narrow-dark::before{
    background: url("@/assets/forest_dark.png");
}
.form-narrow::before{
    background: url("@/assets/forest.png");
}
.form-narrow::before , .form-narrow-dark::before{
    content: '';
    position: absolute;
    width: 100%;
    height: 256px;
    background-size: cover;
    background-repeat: no-repeat;  
}

.input-group input{
    margin-bottom: var(--font-xs);
}
.input-group-error input{
    margin-bottom: 0;
    border-style: solid;
    border: 1px solid var(--error-color)
}
.input-group-error input:focus{
    border-style: none;
}
.label{
    color:var(--main-color);
    font-size:var(--font-xs);
    margin-bottom: 1vh;
    font-weight: bold;
}
.input-group-error label{
    color: var(--error-color);
}
.err-msg{
    font-size: var(--font-xs);
    color:var(--error-color);
}

#form-post-content{
    height:20vh;
    margin-bottom: 2vh;
    resize: none;
}

#form-narrow-one{
    margin-left: 0.5vw;
    flex-grow:1;
    z-index:5;
}

.fox-drawing{
   background-image: url("@/assets/fox-relax.png");
}
.fox-drawing-dark{
   background-image: url("@/assets/fox-relax_dark.png");
}
.fox-drawing, .fox-drawing-dark{
    margin-top:5vh;
    background-repeat: no-repeat;
    background-size: contain;
    flex-grow:4;
    height: 256px; 
    z-index:5;
}
@media (max-width: 1250px){
    .form-narrow, .form-narrow-dark{
        display: block;
        
    }
    #form-narrow-one{position:relative;}
    .fox-drawing, .fox-drawing-dark{
        margin-top:4vh;
        position:relative;
        margin-left: 50%;
        transform: translateX(-50%) 
    }
}

@media (max-width: 900px){
    #form-narrow-two{
        display: none;
    }
    .form-narrow::before , .form-narrow-dark::before{
        display: none;
    }
}
@media (max-height: 550px){
    #form-narrow-two{
        display:none; 
    }
    .form-narrow::before , .form-narrow-dark::before{
        display: none;
    }
}



</style>