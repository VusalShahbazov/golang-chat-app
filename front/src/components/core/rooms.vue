<template>
<div>
  <v-dialog v-model="roomDialog" width="500px">
    <v-card>
      <v-toolbar dark class="primary">
        <v-toolbar-title>
          Create Room
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn @click="closeRoomDialog" class="primary" icon>
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>
      <v-card-text class="mt-4 pb-0">
        <v-col cols="12">
          <v-text-field
              name="name"
              outlined
              @keypress.enter.prevent="saveRoom"
              :rules="[required]"
              v-model="selectedRoomForUpdateOrCreate.name"
              label="Room Name"
          >
          </v-text-field>
        </v-col>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn :loading="loadingRoom" @click="saveRoom" :disabled="!selectedRoomForUpdateOrCreate.name"
               class="primary">Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

    <v-list-item>
      <v-list-item-content>
        <v-list-item-title>Active Rooms</v-list-item-title>
        <v-list-item-subtitle>Select Room and start Message or create your own</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-action>
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
                class="primary"
                dark
                icon
                v-bind="attrs"
                v-on="on"
                @click="roomDialog = true"
            >
              <v-icon>mdi-plus</v-icon>
            </v-btn>
          </template>
          <span>Create Room</span>
        </v-tooltip>
      </v-list-item-action>
    </v-list-item>
    <v-card-actions>
      <v-divider class=""></v-divider>
      <span class="mx-4 mb-2">rooms</span>
      <v-divider class=""></v-divider>
    </v-card-actions>
    <v-card flat class="pb-3" :class="{'scrollable' : $vuetify.breakpoint.mdAndUp}">
      <template v-for="(room  ,index ) in rooms">
        <v-divider v-if="index !== 0" :key="index" class="mx-2"/>
        <v-list-item
            :class="{'v-list-item--active' : room.id === selectedRoom.id}"
            @click.prevent="selectRoom(room)" :key="room.d">
          <v-list-item-content>
            <v-list-item-title class="black--text">
              {{ room.name }}
              <v-chip small color="primary" v-if="room.countofmessages">{{ room.countofmessages }}</v-chip>
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ room.owner.first_name }}
              {{ room.owner.last_name }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-card>
</div>
</template>

<script>
import validate from "../../mixins/validate";
export default {
  mixins: [validate],
  data(){
      return {
        selectedRoomForUpdateOrCreate: {
          name: ""
        },
        loadingRoom: false,
        roomDialog: false,

      }
  },
  methods:{
    closeRoomDialog() {
      this.selectedRoomForUpdateOrCreate = {
        name: ''
      }
      this.roomDialog = false
    },
    loadRoomData(id) {
      this.$axios.get(`api/rooms/${id}`)
          .then(res => {
            this.$store.dispatch("data/SELECT_ROOM" , res.data)
                .then(()=>{
                  this.$nextTick(()=>{
                    var container = document.querySelector("#chatBoxId");
                    container.scrollTop = container.scrollHeight
                  })
                })
                .catch(err => {
                  this.handleError(err)
                })
          })
          .catch((err) => {
            this.handleError(err)
          })
    },
    saveRoom() {
      if (!this.selectedRoomForUpdateOrCreate.name) {
        this.$store.dispatch("snackbar/SET_DATA", {
          show: true,
          color: "info",
          text: "Enter Room Name"
        })
        return
      }

      this.loadingRoom = true
      if (!this.selectedRoomForUpdateOrCreate.id) {
        this.$axios.post('api/rooms', this.selectedRoomForUpdateOrCreate)
            .then(() => {
              this.$store.dispatch("snackbar/SET_DATA", {
                show: true,
                color: "success",
                text: "Room Create"
              })
              this.$store.dispatch("data/LOAD_ROOMS")
            })
            .catch(err => {
              this.handleError(err)
            })
            .finally(() => {
              this.closeRoomDialog()
              this.loadingRoom = false
            })
      } else {
        this.$axios.put(`api/rooms/${this.selectedRoomForUpdateOrCreate.id}`)
            .then(() => {
              this.$store.dispatch("snackbar/SET_DATA", {
                show: true,
                color: "success",
                text: "Successfully updated"
              })
              this.$store.dispatch("data/LOAD_ROOMS")
            })
            .catch(err => {
              this.handleError(err)
            })
            .finally(() => {
              this.closeRoomDialog()
              this.loadingRoom = false
            })
      }
    },
    selectRoom(room){
        this.$store.dispatch("data/SELECT_ROOM" , room)
            .then(()=>{
              this.$nextTick(()=>{
                var container = document.querySelector("#chatBoxId");
                container.scrollTop = container.scrollHeight
              })
            })
            .catch(err => {
              this.handleError(err)
            })

        if (this.selectedRoom.id && this.$route.query.room_id !== this.selectedRoom.id.toString())
          this.$router.replace({query: {room_id: room.id}})
    }
  },
  created() {
    if (this.$route.query.room_id) {
      this.loadRoomData(this.$route.query.room_id)
    }
  },
  computed: {
    rooms() {
      return this.$store.getters['data/rooms']
    },
    selectedRoom(){
      return this.$store.getters['data/selectedRoom']
    }

  },
  name: "rooms"
}
</script>

<style scoped>

</style>