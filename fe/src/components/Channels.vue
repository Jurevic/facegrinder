
<template>
  <div id="topDiv">
    <v-navigation-drawer right absolute temporary v-model="drawer">
      <v-list class="pa-1">
        <v-list-tile>
          <v-list-tile-content>
            <v-list-tile-title><h2>Select channel</h2></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
      <v-list class="pa-1">
        <v-divider></v-divider>
        <v-list-tile v-for="item in items" :key="item.name" @click="">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.name }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <div align="center">
          <v-dialog v-model="dialog" max-width="500px">
            <v-btn color="primary" dark slot="activator">Add channel</v-btn>
            <v-card>
              <v-card-title>
                <span class="headline">New channel</span>
              </v-card-title>
              <v-card-text>
                <v-container grid-list-md>
                  <v-layout wrap>
                    <v-flex xs18 sm9 md6>
                      <v-text-field label="Channel name" v-model="editedItem.name"></v-text-field>
                    </v-flex>
                    <v-flex xs18 sm9 md6>
                      <v-text-field
                        label="Key"
                        hint="At least 8 characters"
                        min="8"
                        :append-icon="keyVisible ? 'visibility' : 'visibility_off'"
                        :append-icon-cb="() => (keyVisible = !keyVisible)"
                        :type="keyVisible ? 'text' : 'password'"
                        v-model="editedItem.key"
                      ></v-text-field>
                    </v-flex>
                  </v-layout>
                </v-container>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" flat @click.native="close">Cancel</v-btn>
                <v-btn color="blue darken-1" flat @click.native="save">Save</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </div>
      </v-list>
    </v-navigation-drawer>
    <v-card>
      <v-card-text>
        <h1>Channels</h1>
      </v-card-text>
    </v-card>
    <video-player
      id="player"
      ref="videoPlayer"
      class="video-player-box"
      :options="playerOptions"
      :playsinline="true">
    </video-player>
    <v-btn
      v-if="!drawer"
      @click.stop="drawer = !drawer"
      color="primary"
      dark slot="activator"
      fixed
      fab
      large
      bottom right
      class="mx-2 my-5"
    ><v-icon>chevron_left</v-icon></v-btn>
  </div>
</template>


<script>
  import axios from 'axios'

  export default {
    data () {
      return {
        drawer: false,
        dialog: false,
        keyVisible: false,
        playerOptions: {
          muted: true,
          language: 'en',
          techOrder: ['flash'],
          sources: [{
            type: 'rtmp/flv',
            src: 'rtmp://localhost/1'
          }]
        },
        items: [],
        editedIndex: -1,
        editedItem: {
          name: '',
          key: ''
        },
        defaultItem: {
          name: '',
          key: ''
        }
      }
    },

    mounted () {
      this.adjustPlayerSize()
      window.addEventListener('resize', this.resizeEvent)
    },

    updated () {
      this.adjustPlayerSize()
    },

    created () {
      this.initialize()
    },

    computed: {
      player () {
        return this.$refs.videoPlayer.player
      }
    },

    beforeDestroy: function () {
      window.removeEventListener('resize', this.resizeEvent)
    },

    methods: {
      resizeEvent (event) {
        this.adjustPlayerSize()
      },

      adjustPlayerSize () {
        let topDiv = document.getElementById('topDiv')

        this.player.height(window.innerHeight - 174)
        this.player.width(topDiv.clientWidth)
      },

      initialize () {
        axios.get(this.$store.state.endpoints.channels, {
          headers: {
            Authorization: this.$store.state.jwt
          }
        }).then((response) => {
          this.items = response.data
        }).catch((error) => {
          console.log(error)
        })
      },

      editItem (item) {
        this.editedIndex = this.items.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog = true
      },

      deleteItem (item) {
        const index = this.items.indexOf(item)
        confirm('Are you sure you want to delete this item?') && this.items.splice(index, 1)
      },

      close () {
        this.dialog = false
        setTimeout(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        }, 300)
      },

      save () {
        axios.post(this.$store.state.endpoints.channels, this.editedItem, {
          headers: {
            Authorization: this.$store.state.jwt
          }
        }).then((response) => {
          if (this.editedIndex > -1) {
            Object.assign(this.items[this.editedIndex], response.data)
          } else {
            this.items.push(response.data)
          }
          this.close()
        }).catch((error) => {
          console.log(error)
        })
      }
    }
  }
</script>
