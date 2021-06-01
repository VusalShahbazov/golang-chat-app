import store from "../store"
import router from "../router"

export function auth(to, from, next) {
    if (!store.getters["auth/loggedIn"]) {
        router.push({name:"account-login"})
    }
    return next();
}
export function noauth(to, from, next) {
    if (store.getters["auth/loggedIn"]) {
        router.push("/")
    }
    return next();
}