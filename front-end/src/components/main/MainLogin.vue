<template>
    <main class="main-wrapper">
        <div class="main-container">

           <div class="login-wrapper">
            <form class="login-form">

                <p class="login-title">Log in</p>

                <div class="login-inputs">
                    <div :class="v$.form.username.$error? errorClass : activeClass">
                        <label for="login-username" class="label">E-mail or username</label>
                        <input v-model="v$.form.username.$model"  autocomplete="username" id="login-username"  class="input-field" @input="clearError()" @blur="v$.form.username.$touch" >
                        <span v-if="v$.form.username.$error" class="err-msg">
                                {{ v$.form.username.$errors[0].$message }}
                        </span>
                    </div>

                    <div :class="v$.form.password.$error? errorClass : activeClass" id="password-container">
                        <label for="login-password" class="label">Password</label>
                        <input v-model="v$.form.password.$model" id="login-password"  autocomplete="current-password" class="input-field" @blur="v$.form.password.$touch" type="password" @keyup.enter.exact="login" @input="clearError()">
                        <span id="show-btn" class="pwd-hidden" @click="tooglePwdVisibility"></span>
                        <span v-if="v$.form.password.$error" class="err-msg">
                                {{ v$.form.password.$errors[0].$message }}
                        </span>
                    </div>
                
                </div>

                <div class="login-checkbox-container" :class="v$.form.checkbox.$error? errorClass : activeClass">
                    <div class="checkbox-interaction">

                        <label for="rum-checkbox" class="label-checkbox" id="rum-checkbox-label">Check the box if you just took a sip of rum</label>
                        <input v-model="v$.form.checkbox.$model" id="rum-checkbox" type="checkbox">

                    </div>
                    <span v-if="v$.form.checkbox.$error" class="err-msg" id="checkbox-error">
                            {{ v$.form.checkbox.$errors[0].$message }}
                    </span>

                </div>
                <div @mouseup="login" id="login-btn" class="btn">Log in</div>
            </form>   
                <login-graphics></login-graphics>
            </div> 
        
        </div>
    </main>
</template>
<script>
import LoginGraphics from "@/components/main/login/LoginGraphics.vue"
import {mapWritableState} from 'pinia'
import {useMainStore} from '@/stores/main.js'
import axios from 'axios'
/* ------------------------------- Validation ------------------------------- */
import useValidate from "@vuelidate/core";
import { required, helpers, sameAs} from "@vuelidate/validators";


export default {
    components:{LoginGraphics},
    props:["resp"],
    data: function(){
        return{
            serverUrl: "ws://localhost:8081/ws",
            serverErrors : {},
            v$: useValidate(),
            form:{
                username:"",
                password:"", 
                checkbox:"" 
            }, 
            activeClass: "input-group",
            errorClass: "input-group-error"
        }
    },
    computed:{
        ...mapWritableState(useMainStore,["ws"]),
        ...mapWritableState(useMainStore,["isLoggedIn"]),
        ...mapWritableState(useMainStore,["name"]),
    },
    methods:{
        // ...mapActions(useMainStore, {resetStore : 'reset'}),
        async login(){
            const isFormCorrect = await this.v$.$validate()
            if (!isFormCorrect) return
            this.serverErrors = {}; //clean server error list
            axios.post("http://localhost:8081/api/login", this.form,{ withCredentials: true })
            .then((res)=> {
                if (res.data.length != 0){
                    res.data.forEach(err => {
                        this.serverErrors[err.property] = err.message
                    });
                    this.v$.$reset();
                    this.v$.$touch();
                }else{
                    // Login success
                    this.isLoggedIn = true;
                    // this.name = this.form.username;
                    this.connectToWebsocket();
                }
            })
        },
        goToFeed() {
            this.$router.push({name:'feed'})
        },
        connectToWebsocket(){
            this.ws = new WebSocket( this.serverUrl);
            this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
            this.ws.addEventListener('error', (event) => { this.onWebsocketError(event) });
            this.ws.addEventListener('close', (event) => { this.onWebsocketClose(event) });
        },
        onWebsocketOpen() {
             this.ws.send(JSON.stringify({
                 action:'get-username'
             }))
            this.goToFeed();
        },
        onWebsocketError(event) { 
            console.log("Oh..... Error :  ", event)
        },
        onWebsocketClose(event) { 
            console.log("Connection closed:  ", event)
            // console.log("Reset storage")
            // this.resetStore()
            //  localStorage.removeItem('main')
        },
        tooglePwdVisibility(){
            const pwdField = document.querySelector("#login-password");
            const btn = document.querySelector("#show-btn");
            if(pwdField.type == "password"){
                pwdField.type  = "text"
                btn.classList.remove("pwd-hidden")
                btn.classList.add("pwd-visible")
            }else{
                pwdField.type = "password"
                btn.classList.remove("pwd-visible")
                btn.classList.add("pwd-hidden")
            }
        },
        clearError(){
            delete this.serverErrors['username'];
        }
    },
    validations(){
        const that = this;
        return{
            form:{
                username:{
                    required : helpers.withMessage("Fox without a name or adress is lost in the forest", required),
                    serverValidation:{
                        $message() {
                            return that.serverErrors.username;
                        },
                        $validator() {
                            return !that.serverErrors.username;
                        }
                    }
                },
                password:{
                    required: helpers.withMessage("Password is required", required),
                    serverValidation:{
                        $message() {
                            return that.serverErrors.password;
                        },
                        $validator() {
                            return !that.serverErrors.password;
                        }
                    }
                },
                 checkbox:{
                     sameAs:helpers.withMessage("Please take a sip and check that box",sameAs(true)) 
                 }
            }
        }
    }
}
</script>

<style scoped>
.login-wrapper{
    height:100%;
    display:flex;
}
.login-form{
    padding-right: 5vh;
    padding-left: 1vh;
    width: 36vw;
    overflow-y: auto; 
}
@media (max-width: 900px){
    .login-form{
        width: 100%;
        padding-right: 1vh;
    }
}
.signup-form *{
   display: block; 
}
.login-title{
    font-weight: bold;
    margin-bottom: 2vh;
}
.login-inputs{
    text-align: left;
}
.input-group-error input{
    margin-bottom: 0;
    border-style: solid;
    border: 1px solid var(--error-color)
}
.input-group input{
    margin-bottom: var(--font-xs);
}
.input-group-error input:focus{
    border-style: none;
}
.login-checkbox-container{
    text-align: right;
    position: relative;
    margin-top: 2vh;
}
.checkbox-interaction{
    display: flex;
    flex-wrap: nowrap;
    justify-content:flex-end;
}
.label{
    color:var(--main-color);
    font-size:var(--font-s);
    font-weight: bold;
}
.input-group-error label{
    color: var(--error-color);
}
.err-msg{
    font-size: var(--font-xs);
    color:var(--error-color);
}
#login-btn{
    width: 10vw;
    margin-left: auto;
    margin-top: 5vh;
}

/* checkbox horror */
.label-checkbox{
    color:var(--main-color);
    font-weight: lighter;
    font-size:var(--font-xs);
    margin-right:25px;
}
#rum-checkbox{
    visibility:hidden;
     -moz-appearance:initial;
}
#rum-checkbox-label{
    cursor: pointer;
    padding: 0;
}
#rum-checkbox:before{
    content: '';
    position: absolute;
    right: 0;
    width: 15px;
    height: 15px;
    border-radius: 15%;
    background: var(--content-color-active);
    visibility:visible;

}
#rum-checkbox:hover:before{
    background: var(--accent-color-dark);

}
#rum-checkbox:focus:before{
    box-shadow: 0 0 0 3px rgba(0, 0, 0, 0.12);
}
#rum-checkbox:checked:before{
    background: var(--accent-color-light);
}
#rum-checkbox:checked:after{
    content: '';
    position: absolute;
    right: 1px;;
    width: 12px;
    height: 15px;
    background-image: url("@/assets/icons/check.svg");
    background-size: contain;
    background-repeat: no-repeat;
    visibility:visible;
    filter: invert(74%) sepia(89%) saturate(1%) hue-rotate(2deg) brightness(109%) contrast(100%);
}
#checkbox-error{
    display:block;
}
/* -------------------------- show password toogle -------------------------- */
#password-container{
    position: relative;
}
#show-btn{
   position: absolute;
   top: 28px;
   cursor:pointer;
    background-repeat: no-repeat;
   filter: invert(84%) sepia(6%) saturate(240%) hue-rotate(101deg) brightness(85%) contrast(91%);
}
.pwd-hidden{
    background-image: url("@/assets/icons/see.svg");
    height: 14px;
    width: 14px;
    right: 11px;
}
.pwd-visible{
    background-image: url("@/assets/icons/hide.svg");
    height: 16px;
   width: 16px;
   right: 10px;;
}
</style>