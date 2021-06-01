/* global axios */

export default {
    namespaced: true,

    state: () => ({
        user: null,
        status: '',
        token: localStorage.getItem("token") || "",
    }),

    getters: {
        status: state => state.status,
        user: state => state.user,
        loggedIn: state => state.status === 'success',
        token: state => state.token,
    },

    mutations: {
        auth_request(state) {
            state.status = "loading";
        },

        auth_success(state) {
            state.status = "success";
        },

        auth_error(state) {
            state.status = "error";
        },
        logout(state) {
            state.status = "";
            state.token = "";
            state.user = null;
        },
        auth_set_user(state, user) {
            state.user = user;
        },

        set_token(state, token) {
            state.token = token
        }
    },

    actions: {
        LOGIN(context, payload) {
            return new Promise((resolve, reject) => {
                context.commit("auth_request");
                axios
                    .post("api/auth/login", payload)
                    .then(res => {
                        localStorage.setItem("token", res.data.token);
                        context.commit("auth_success");
                        context.commit("set_token", res.data.token);
                        context.commit("auth_set_user", res.data.user);
                        resolve(res);
                    })
                    .catch(err => {
                        context.commit("auth_error");
                        reject(err);
                    });
            });
        },
        REGISTER(context, payload) {
            return new Promise((resolve, reject) => {
                context.commit("auth_request");
                axios
                    .post("api/auth/signup", payload)
                    .then(res => {
                        localStorage.setItem("token", res.data.token);
                        context.commit("auth_success");
                        context.commit("set_token", res.data.token);
                        context.commit("auth_set_user", res.data.user);
                        resolve(res);
                    })
                    .catch(err => {
                        context.commit("auth_error");
                        reject(err);
                    });
            });
        },
        LOGOUT(context) {
            return new Promise((resolve, reject) => {
                axios
                    .post("api/auth/logout")
                    .then(res => {
                        context.commit("logout");
                        localStorage.removeItem("token")
                        resolve(res);
                    })
                    .catch(err => {
                        reject(err);
                    });
            });
        },
        GET_PROFILE(context) {
            return new Promise((resolve, reject) => {
                if (context.state.token) {
                    context.commit("auth_request")

                    axios
                        .get("api/auth/me")
                        .then(res => {
                            context.commit("auth_success")
                            context.commit("auth_set_user", res.data.user);
                            resolve(res);
                        })
                        .catch(err => {
                            context.commit("auth_error")
                            reject(err);
                        });
                } else
                    reject("neet to logged in")
            });
        },
    }
}