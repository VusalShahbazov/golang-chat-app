import store from "../store/index";
import router from "../router";

export default {
    data() {
        return {
            errors: {},
            errorMessageFromRequest: ''
        }
    },
    methods: {
        setErrors(errors) {
            this.errors = errors
        },
        resetError(key) {
            this.errors[key] = undefined
        },

        handleError(error, mainResource = false) {
            if (error.response) {
                const response = error.response
                if (response.status === 422 && response.data.errors) {
                    this.setErrors(response.data.errors)
                    store.dispatch("snackbar/SET_DATA", {
                        color: 'error',
                        text: 'validations.message',
                        show: true
                    })
                    this.errorMessageFromRequest = 'validations.message'

                    return
                }
                if (response.status === 400 && response.data.error) {
                    store.dispatch("snackbar/SET_DATA", {
                        color: 'error',
                        text: response.data.error,
                        show: true
                    })
                    this.errorMessageFromRequest = response.data.error

                    return
                }
                if (response.status === 429) {
                    store.dispatch("snackbar/SET_DATA", {
                        color: 'error',
                        text: "errors.too-many-request",
                        show: true
                    })
                    this.errorMessageFromRequest = "errors.too-many-request"
                    return
                }

                if (response.status.toString().startsWith("5")) {
                    if (mainResource) {
                        router.push({name: 'server-error'})
                    } else {
                        this.unexpectedBehavior(error, true)
                    }
                    return;
                }
            }
            this.unexpectedBehavior(error, false)

        },
        unexpectedBehavior(error, serverError = false) {
            console.error(error)
            this.errorMessageFromRequest = serverError ? 'errors.server-error' : 'errors.unexpected'
            store.dispatch("snackbar/SET_DATA", {
                color: "error",
                show: true,
                text: this.errorMessageFromRequest,
            })
        }
    }
}