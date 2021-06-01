<template>
  <v-container class="d-flex align-center">
    <v-row no-gutters>
      <v-col cols="12" sm="8" md="6" offset-md="3" offset-sm="2">
        <v-form @submit.prevent="register" ref="registerForm">
          <v-card class="pa-5 mt-md-8">
            <v-card-title class="font-weight-bold text-h6 black--text text-center pa-0">
              Sign Up
            </v-card-title>
            <v-card-subtitle class="pl-0 pt-4">
              Есть аккоунт?
              <router-link :to="{name:'account-login'}" class="text-decoration-none">
              <span class="subtitle-2 text-decoration-none">
                Войти
              </span>
              </router-link>
            </v-card-subtitle>
            <v-card-text class="pa-0">
              <v-row no-gutters>
                <v-col cols="12" sm="6" class="pr-sm-2">
                  <v-text-field
                      name="first_name"
                      outlined
                      :rules="[required]"
                      :error-messages="errors['first_name']"
                      @focus="resetError('first_name')"
                      v-model="form.first_name"
                      label="First Name"
                  >
                  </v-text-field>
                </v-col>
                <v-col cols="12" sm="6">
                  <v-text-field
                      name="last_name"
                      outlined
                      :rules="[required]"
                      :error-messages="errors['last_name']"
                      @focus="resetError('last_name')"
                      v-model="form.last_name"
                      label="Last Name"
                  >
                  </v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                      name="email"
                      outlined
                      :rules="[required,email]"
                      :error-messages="errors['email']"
                      @focus="resetError('email')"
                      append-icon="mdi-email"
                      v-model="form.email"
                      label="Email"
                  >
                  </v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                      name="password"
                      outlined
                      :rules="[required , password]"
                      :error-messages="errors['password']"
                      @focus="resetError('password')"
                      v-model="form.password"
                      label="Password"
                      @click:append="showPassword = !showPassword"
                      :type="showPassword ? 'text' : 'password'"
                      :append-icon="showPassword ? 'mdi-lock-open' : 'mdi-lock'"
                  >
                  </v-text-field>
                </v-col>
              </v-row>
            </v-card-text>
            <v-card-actions class="pa-0 ">
              <v-btn
                  :loading="loading" type="submit" class="primary" width="100%" large>
                Register
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import validation from "../../../mixins/validate";

export default {
  name: "Register",
  mixins: [validation],
  computed: {
    url() {
      return process.env.VUE_APP_BASE_URL ? process.env.VUE_APP_BASE_URL : 'http://localhost:8000'
    },
    loading() {
      return this.$store.getters['auth/status'] === 'loading'
    },
  },
  methods: {
    clearTimer() {
      if (this.timerForResendEmail) clearTimeout(this.timerForResendEmail)
      this.canResend = false
    },
    setTimerForResend() {
      this.clearTimer()
      this.canResend = false
      this.timerForResendEmail = setTimeout(() => {
        this.canResend = true
      }, 15000)
    },
    emailOrigin() {
      if (this.form.email) {
        const arr = this.form.email.split("@")
        if (arr.length === 2) {
          return `http://${arr[1]}`
        }
      }
      return "#"
    },
    register() {
      if (!this.$refs.registerForm.validate()) return;

      this.setTimerForResend()
      this.$store.dispatch("auth/REGISTER", this.form)
          .then(() => {
            this.$router.push({name:"messages"})
          })
          .catch(err => {
            this.handleError(err)
          })


    }
  },
  data() {
    return {
      accept: false,
      showPassword: false,
      form: {
        last_name: '',
        first_name: '',
        email: '',
        password: '',
        is_freelancer: false
      },
    }
  },
}
</script>
