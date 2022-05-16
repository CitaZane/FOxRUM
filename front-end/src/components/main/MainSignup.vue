<template>
    <main class="main-wrapper">
        <div class="main-container">

           <div class="signup-wrapper">
            <form class="signup-form" @submit.prevent="signUp" @closed="v$.$reset()" >
                <p class="signup-title">Sign up</p>

                    <!-- /* ------------------------------- first name ------------------------------- */ -->
                    <div :class="v$.form.firstName.$error? errorClass : activeClass">
                        <label for="signup-first-name" class="label">First Name</label>
                        <input v-model="v$.form.firstName.$model" id="signup-first-name"  class="input-field" @blur="v$.form.firstName.$touch" @input="clearError('firstName')">
                        <span v-if="v$.form.firstName.$error" class="err-msg">
                                {{ v$.form.firstName.$errors[0].$message }}
                        </span>
                    </div>
                    <!-- /* -------------------------------- last name ------------------------------- */ -->
                    <div :class="v$.form.lastName.$error? errorClass : activeClass">
                        <label for="signup-last-name" class="label">Last Name</label>
                        <input v-model="v$.form.lastName.$model" id="signup-last-name"  class="input-field" @input="clearError('lastName')">
                        <span v-if="v$.form.lastName.$error" class="err-msg">
                                {{ v$.form.lastName.$errors[0].$message }}
                        </span>
                    </div>

                <div class="form-subcontainer">
                    <!-- /* --------------------------------- gender --------------------------------- */ -->
                     <div :class="v$.form.gender.$error? errorClass : activeClass">
                        <label for="signup-gender" class="label">Gender</label>
                        <select v-model="v$.form.gender.$model" id="signup-gender"  class="input-field" @input="clearError('gender')">
                            <option value="female">Female</option>
                            <option value="male">Male</option>
                            <option value="other">Non-binary</option>
                            <option value="none">Prefer not to say</option>
                        </select>
                        <span v-if="v$.form.gender.$error" class="err-msg">
                                {{ v$.form.gender.$errors[0].$message }}
                        </span>
                    </div>
                    <!-- /* ------------------------------ date of birth ----------------------------- */ -->
                    <div :class="v$.form.birth.$error? errorClass : activeClass">
                        <label for="signup-birth" class="label">Date of birth</label>
                        <input v-model="v$.form.birth.$model" id="signup-birth"  class="input-field" type="date" @input="clearError('birth')">
                        <span v-if="v$.form.birth.$error" class="err-msg">
                                {{ v$.form.birth.$errors[0].$message }}
                        </span>
                    </div>

                </div>
                <!-- /* -------------------------------- username -------------------------------- */ -->
                <div :class="v$.form.username.$error? errorClass : activeClass">
                    <label for="signup-username" class="label">Your Fox Name / Username</label>
                    <input v-model="v$.form.username.$model" id="signup-username"  class="input-field" @input="clearError('username')">
                    <span v-if="v$.form.username.$error" class="err-msg">
                            {{ v$.form.username.$errors[0].$message }}
                    </span>
                </div>
                <!-- /* ---------------------------------- email --------------------------------- */ -->
                <div :class="v$.form.email.$error? errorClass : activeClass">
                    <label for="signup-email" class="label">E-mail</label>
                    <input v-model="v$.form.email.$model" id="signup-email"  class="input-field" type="email" @blur="v$.form.email.$touch" @input="clearError('email')">
                    <span v-if="v$.form.email.$error" class="err-msg">
                            {{ v$.form.email.$errors[0].$message }}
                    </span>
                </div>
                <!-- /* -------------------------------- password -------------------------------- */ -->
                <div :class="v$.form.password.$error? errorClass : activeClass">
                    <label for="signup-password" class="label">Password</label>
                    <input v-model="v$.form.password.$model" id="signup-password"  class="input-field" type="password" @input="clearError('password')">
                    <span v-if="v$.form.password.$error" class="err-msg">
                            {{ v$.form.password.$errors[0].$message }}
                    </span> 
                </div>  
                <!-- /* -------------------------- password confirmation ------------------------- */ -->
                <div :class="v$.form.confirm.$error? errorClass : activeClass">
                    <label for="signup-password-repeat" class="label">Confirm Password</label>
                    <input v-model="v$.form.confirm.$model" id="signup-password-repeat"  class="input-field" type="password"> 
                    <span v-if="v$.form.confirm.$error" class="err-msg">
                            {{ v$.form.confirm.$errors[0].$message }}
                    </span>
                </div>


                <div @mouseup="signUp" id="signup-btn" class="btn">Sign Up</div>
                </form>
                <login-graphics></login-graphics>
            </div> 
        
        </div>
    </main>
</template>
<script>
import axios from 'axios'
import LoginGraphics from "@/components/main/login/LoginGraphics.vue"
import {mapWritableState} from 'pinia'
import {useMainStore} from '@/stores/main.js'
/* ------------------------------- Validation ------------------------------- */
import useValidate from "@vuelidate/core";
import { required, email, minLength, sameAs, alpha, maxLength, helpers} from "@vuelidate/validators";
 /* ---------------------------- Custom validators --------------------------- */
const notBornInFuture = (value)=> Date.parse(value) < new Date().getTime()
const lessThanHundred = (value)=>  new Date(Date.parse(value)).getFullYear() > new Date().getFullYear() - 100
const nameRegex = /^[\w+]+$/
const usernameValid = (value)=> value.match(nameRegex)

export default {
    components:{LoginGraphics},
    props:["resp"],
    data: function(){
        return{
            serverUrl: "ws://localhost:8081/ws",
            serverErrors : {},
            v$: useValidate(),
            form:{
                firstName: "",
                lastName:"",
                gender:"",
                birth: "",
                username:"",
                email:"",
                password: "",
                confirm:""
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
        async signUp(){
            const isFormCorrect = await this.v$.$validate()
            if (!isFormCorrect) return
            this.serverErrors = {}; //clean server error list
            axios.post("http://localhost:8081/api/signup", this.form,{ withCredentials: true })
            .then((res)=> {
                // `res.data` contains the parsed response body
                if (res.data.length != 0){
                    res.data.forEach(err => {
                        this.serverErrors[err.property] = err.message
                    });
                    this.v$.$reset();
                    this.v$.$touch();
                }else{
                    this.isLoggedIn = true
                    this.name = this.form.username;
                    this.connectToWebsocket();
                }
            })
        },
        connectToWebsocket(){
            this.ws = new WebSocket( this.serverUrl);
            this.ws.addEventListener('open', () => { this.onWebsocketOpen() });
            this.ws.addEventListener('error', (event) => { this.onWebsocketError(event) });
            this.ws.addEventListener('close', (event) => { this.onWebsocketClose(event) });
        },
         onWebsocketOpen() {
             console.log("Socket open!"),
            this.goToFeed();
      },
      onWebsocketError(event) { 
            console.log("Oh..... Error :  ", event)
        },
        onWebsocketClose(event) { 
            console.log("Connection closed:  ", event)
            console.log("Delete storage")
            localStorage.removeItem('main')
        },
        goToFeed() {
            this.$router.push({name:'feed'})
        },
        clearError(path){
            delete this.serverErrors[path];
        }
    },
    validations(){
        const that = this;
        return{
            form:{
                firstName: { 
                    required : helpers.withMessage("You surely have a name, and it is required", required), 
                    alpha :helpers.withMessage("Are you an android version? Your name should contain only letters", alpha),
                    maxLength:helpers.withMessage("Really, your name is that long? Keep it shorter so we can remember it!", maxLength(60)),  
                    $lazy: true,
                    serverValidation:{
                        $message() {
                            return that.serverErrors.firstName;
                        },
                        $validator() {
                            return !that.serverErrors.firstName;
                        }
                    }},
                lastName:{ 
                    required: helpers.withMessage("No last name? Sadly it is required", required),
                    // alpha:helpers.withMessage("Feeling creative? Lets stick with letters for your last name", alpha),
                    maxLength:helpers.withMessage("Really, your last name is that long? Keep it shorter so we can remember it!", maxLength(60)),  
                    $lazy: true,
                    serverValidation:{
                        $message() {
                            return that.serverErrors.lastName;
                        },
                        $validator() {
                            return !that.serverErrors.lastName;
                        }
                    }
                    },
                gender:{ 
                    required , 
                    alpha,
                    serverValidation:{
                        $message() {
                            return that.serverErrors.gender;
                        },
                        $validator() {
                            return !that.serverErrors.gender;
                        }
                    }},
                birth: { required : helpers.withMessage("Date of birth is required", required),
                notBornInFuture: helpers.withMessage("Are you a time traveler? Cool", notBornInFuture),
                lessThanHundred: helpers.withMessage("Oh, really more than 100 years old?", lessThanHundred),
                $lazy: true,
                serverValidation:{
                        $message() {
                            return that.serverErrors.birth;
                        },
                        $validator() {
                            return !that.serverErrors.birth;
                        }
                    }
                },
                username:{ 
                    required : helpers.withMessage("Fox without a name is lost in the forest", required), 
                    minLength: helpers.withMessage("Your name is a bit too short. Min 4 characters", minLength(4)),
                    maxLength:helpers.withMessage("Max 15 characters. There are some standards even in forest.", maxLength(15)),
                    usernameValid: helpers.withMessage("Sorry, the limit is letters, numbers and  underscore", usernameValid),
                    $lazy: true,
                    serverValidation:{
                        $message() {
                            return that.serverErrors.username;
                        },
                        $validator() {
                            return !that.serverErrors.username;
                        }
                    }
                },
                email:{ 
                    required: helpers.withMessage("Email is required. Even for foxes", required), 
                    email: helpers.withMessage("Somehow this doesn't seem as a valid email...", email),
                    $lazy: true,
                    serverValidation:{
                        $message() {
                            return that.serverErrors.email;
                        },
                        $validator() {
                            return !that.serverErrors.email;
                        }
                    }
                    },
                password: { 
                    required: helpers.withMessage("Password is required", required)
                    , minLength:helpers.withMessage("For minimal safety, lets use at least 6 characters", minLength(6)),
                    $lazy: true,
                    serverValidation:{
                        $message() {
                            return that.serverErrors.password;
                        },
                        $validator() {
                            return !that.serverErrors.password;
                        }
                    }
                    },
                confirm:{ 
                    required: helpers.withMessage("Confirmation password is required", required),
                    sameAs: helpers.withMessage("Passwords are like twins. They should match", sameAs(this.form.password)),
                    $lazy: true }
            }
        }
    }
}
</script>


<style scoped>
.signup-wrapper{
    height:100%;
    display:flex;
}
.signup-form{
    padding-right: 5vh;
    padding-left: 1vh;
    width: 36vw;
    overflow-y: auto;
    
}
@media (max-width: 900px){
    .signup-form{
        width: 100%;
        padding-right: 1vh;
    }
}
.signup-form *{
   display: block;
    
}
.signup-title{
    font-weight: bold;
    margin-bottom: 2vh;
    margin-left: auto;
    margin-right: auto;
}
.input-group{
    text-align: left;
}
.input-group input{
    margin-bottom: var(--font-xs);
}
.input-group-error{
    text-align: left;
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
    font-size:var(--font-s);
    font-weight:bold;
}
.input-group-error label{
    color: var(--error-color);
}
.form-subcontainer{
    width: 100%;
    display:flex;
    flex-direction: column;
    flex-wrap: wrap;
    height: 50px;
    column-gap: 8%;
    margin-bottom: var(--font-xs)
}
.form-subcontainer div{
flex-grow: 1;
    width: 46%;
}

#signup-btn{
    width: 10vw;
    margin-left: auto;
    margin-top: 5vh;
}
.err-msg{
    font-size: var(--font-xs);
    color:var(--error-color);
}
</style>