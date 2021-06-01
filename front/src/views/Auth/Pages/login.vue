<template>
  <v-container class="d-flex align-center">
    <v-row no-gutters>
      <v-col cols="12" sm="6" md="4" offset-md="4" offset-sm="3">
        <v-form ref="loginFrom" @submit.prevent="login">
          <v-card class="mt-md-8  pa-5">
            <v-card-title class="font-weight-bold text-h6 black--text text-center pa-0">
              Sign in to your account
            </v-card-title>
            <v-card-subtitle class="pl-0 pt-4">
              Новый пользователь?
              <router-link :to="{name:'account-register'}" class="text-decoration-none">
              <span class="subtitle-2 text-decoration-none">
                Создать учетную
              </span>
              </router-link>
            </v-card-subtitle>
            <v-card-text class="pa-0">
              <v-row no-gutters>
                <v-col cols="12">
                  <v-text-field
                      name="email"
                      outlined
                      :rules="[required , email]"
                      v-model="form.email"
                      append-icon="mdi-email"
                      label="Email"
                  >
                  </v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                      name="password"
                      outlined
                      :rules="[required , min6]"
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

              <v-btn :loading="loading" type="submit" class="primary" width="100%" large>
                Login
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import validation from "../../../mixins/validate"

export default {
  name: "Login",
  mixins: [validation],
  data() {
    return {
      form: {
        email: '',
        password: ''
      },
      showPassword: false
    }
  },
  created() {

  },
  methods: {
    login() {
      if (!this.$refs.loginFrom.validate()) return
      this.$store.dispatch("auth/LOGIN", this.form)
          .then(() => {
            this.$store.dispatch("snackbar/SET_DATA", {
              text: "Successfull logged in",
              color: "success",
              show: true
            })
            this.$router.push("/")
          })
          .catch(err => {
            this.handleError(err)

          })
    }
  },
  computed: {
    loading() {
      return this.$store.getters['auth/status'] === 'loading'
    },
  }
}
</script>

