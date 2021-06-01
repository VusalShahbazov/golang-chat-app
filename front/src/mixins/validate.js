import {required, email, maxLength, minLength} from "vuelidate/lib/validators"
import response from "./response";
export default {
    data() {
        return {
            errors: {},
            errorMessageFromRequest: ''
        }
    },
    mixins:[response],
    methods: {
        required(value) {
            return required(value) || "validations.required"
        },
        email(value) {
            return email(value) || "validations.email"
        },
        checked(value) {
            return !!value || "validations.checked"
        },
        max50(value) {
            return maxLength(50)(value) || "validations.max50"
        },
        max100(value) {
            return maxLength(100)(value) || "validations.max50"
        },
        max200(value) {
            return maxLength(200)(value) || "validations.max200"
        },
        min6(value) {
            return minLength(6)(value) || "validations.min6"
        },
        password(value) {
            return /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{8,}$/.test(value) ||
                "validations.password"
        },
        same(value1, value2) {
            return (value1 === value2) || "validations.not-same"
        },

    }
}