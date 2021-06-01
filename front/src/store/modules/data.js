/* global axios */

import Vue from "vue";

export default {
    namespaced: true,

    state: () => ({
        rooms: [],
        selectedRoom:{},
        messages: [],
        loading_messages:false
    }),

    getters: {
        rooms: state => state.rooms,
        selectedRoom: state => state.selectedRoom,
        messages : state => state.messages,
        loading_messages : state => state.loading_messages,
    },

    mutations: {
        set_room(state, room) {
            state.selectedRoom = room
            let rm = state.rooms.find(el => el.id === room.id)
            if (rm) {
                let index = state.rooms.indexOf(rm)
                rm.countofmessages = 0
                Vue.set(state.rooms, index, rm)
            }
        },
        loading_messages(state ,bool){
            state.loading_messages = bool
        },
        set_messages(state ,messages){
            state.messages = messages

        },
        push_message(state , message){
            state.messages.push(message)
        },
        set_rooms(state, rooms) {
            state.rooms = rooms
        },
        set_room_msm(state, id) {
            let room = state.rooms.find(el => el.id === id)
            if (room) {
                let index = state.rooms.indexOf(room)
                room.countofmessages = 1
                Vue.set(state.rooms, index, room)
            }
        },
    },

    actions: {
        SELECT_ROOM(context , payload){
            return new Promise((resolve, reject) => {
                context.commit("set_room" , payload)
                context.commit("loading_messages" , true)
                this.loadingMessages = true;
                axios.get(`api/rooms/${payload.id}/messages?page=1&per_page=25`)
                    .then(res => {
                        context.commit("set_messages" , res.data)
                        resolve()
                    })
                    .catch((err) => {
                        reject(err)
                    })
                    .finally(() => {
                        context.commit("loading_messages" , false)
                    })
            })

        },
        LOAD_ROOMS(context) {
            return new Promise((resolve, reject) => {
                axios.get(`api/rooms`)
                    .then(res => {
                        context.commit('set_rooms', res.data)
                        resolve(res)
                    })
                    .catch(err => {
                        reject(err)
                    })
            });
        },
    }
}