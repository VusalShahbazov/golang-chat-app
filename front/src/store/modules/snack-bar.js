export default {
    namespaced:true,
    state: () => ({
        snack: {
            show: false,
            text: '',
            // [success , error , warning , info]
            color: 'info',
            timeout: 6000,
        },
    }),

    getters: {
        snack: state => state.snack,
    },
    mutations: {
        set(state, data) {
            if (data.color)  state.snack.color = data.color
            state.snack.timeout = data.timeout ? data.timeout : 6000
            state.snack.text = data.text
            state.snack.show = data.show
        },
    },
    actions: {
        SET_DATA(context, payload) {
            context.commit("set", payload)
        },
        HIDE_SNACK_BAR(context){
            context.commit("set", {
                show : false,
                text: '',
            })
        }
    }
}