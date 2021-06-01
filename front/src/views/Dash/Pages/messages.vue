<template>
  <v-container>

    <v-row>
      <v-col cols="12" md="4" v-if="$vuetify.breakpoint.mdAndUp">
        <v-card flat>
          <rooms/>
        </v-card>
      </v-col>
      <v-col cols="12" md="8">
        <v-card flat>
          <v-toolbar elevation="0">
            <v-toolbar-title class="d-flex align-center">
              {{ selectedRoom.name ? selectedRoom.name : 'Select Room' }}
            </v-toolbar-title>
          </v-toolbar>
          <v-card flat class="scrollable" id="chatBoxId" ref="chatBox">
            <v-card-text v-if="!loadingMessages">
              <div v-if="!messagesOrder.length" class="d-flex justify-center">
                <v-card flat>
                  <v-card-text class="text-center">
                    <v-icon color="primary" size="50">{{
                        selectedRoom.id ? 'mdi-forum' : 'mdi-cursor-default-click'
                      }}
                    </v-icon>
                  </v-card-text>
                  <v-card-title>{{ selectedRoom.id ? 'No Message' : 'Select Room' }}</v-card-title>
                </v-card>
              </div>
              <div v-else>
                <div :key="message.id" v-for="message in messagesOrder">
                  <div v-if="message.from_user.id === user.id" class=" text-right  mb-5">
                    <span class="my px-6 py-2 primary white--text ">
                    {{ message.body }}
                    </span>
                  </div>
                  <div class="d-inline-block notmy" v-else>
                    <span class="font-weight-bold">
                    {{ message.from_user.first_name }}
                    {{ message.from_user.last_name }}
                    </span>
                    <br>
                    <span>
                    {{ message.body }}

                    </span>
                  </div>
                </div>
              </div>
            </v-card-text>
            <loader v-else/>
          </v-card>

          <v-card-actions>

            <v-text-field
                outlined
                hide-details
                @keypress.enter.prevent="sendSms"
                v-model="sms"
            >

            </v-text-field>
            <v-btn @click="sendSms" large rounded class="primary mx-3">
              <v-icon>mdi-send</v-icon>
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import validate from "../../../mixins/validate";
import Loader from "../../../components/core/loader";
import Rooms from "../../../components/core/rooms";

export default {
  components: {Rooms, Loader},
  mixins: [validate],
  name: "messages",
  data() {
    return {
      sms: '',
      ws: null,
      messages: [],
      page: 1,
    }
  },
  methods: {
    ScrollToMessages() {
      if (this.$refs.chatBox)
        this.$refs.chatBox.$el.scrollTop = this.$refs.chatBox.$el.scrollHeight
    },

    sendSms() {
      if (!this.sms) return
      this.ws.send(JSON.stringify({room_id: this.selectedRoom.id, body: this.sms}))
      this.sms = ''
    },
    startAndConfigureWs() {
      this.ws = new WebSocket(`ws://localhost:3000/ws?token=${this.$store.getters['auth/token']}`)
      this.ws.onmessage = (params) => {
        let data = JSON.parse(params.data)
        if (data.for_message) {
          if (this.selectedRoom.id === data.message.room_id) {
            this.$store.commit("data/push_message" ,data.message )
            this.$nextTick(() => {
              this.ScrollToMessages()
            })
          } else {
            let room = this.rooms.find(el => el.id === data.message.room_id)
            if (!room) {
              this.$store.dispatch("data/LOAD_ROOMS")
                  .then(() => {
                    this.$store.commit("data/set_room_msm", data.message.room_id)
                  })
            } else {
              if (room.id !== this.selectedRoom.id)
                this.$store.commit("data/set_room_msm", room.id)
            }
          }
        }
      }
      this.ws.onclose = (evt) => {
        console.log(evt)
        this.$store.dispatch("snackbar/SET_DATA", {
          show: true,
          text: "Connection clothed",
          color: "error"
        })
      }
      this.ws.onconnetcion = (evt) => {
        console.log(evt)
        this.$store.dispatch("snackbar/SET_DATA", {
          show: true,
          text: "Connected",
          color: "success"
        })
      }
    }
  },
  created() {
    this.startAndConfigureWs()
    this.$store.dispatch("data/LOAD_ROOMS")

  },
  computed: {
    rooms() {
      return this.$store.getters['data/rooms']
    },
    user() {
      return this.$store.getters['auth/user']
    },
    messagesOrder() {
      // eslint-disable-next-line vue/no-side-effects-in-computed-properties
      return this.$store.getters['data/messages'].sort((a, b) => a.id - b.id)
    },
    selectedRoom(){
      return this.$store.getters['data/selectedRoom']
    },
    loadingMessages(){
      return this.$store.getters['data/loading_messages']
    }
  },
}
</script>

<style>
.scrollable {
  height: calc(100vh - 220px);
  overflow-y: scroll;
  max-height: calc(100vh - 200px)
}

.my {
  border-radius: 15px 15px 0 15px;
}

.notmy {
  text-align: left;
  border-radius: 15px 15px 15px 0;
  border: 1px solid #888888;
  padding: 5px 10px;
  margin-bottom: 10px;

}
</style>