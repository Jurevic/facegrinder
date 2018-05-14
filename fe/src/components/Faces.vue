<template>
  <div>
    <v-card>
      <v-card-text>
        <h1>Faces</h1>
      </v-card-text>
    </v-card>
    <v-container fluid grid-list-md>
      <v-layout row wrap>
        <v-flex v-bind="{ [`xs2`]: true }" v-for="item in items" :key="item.name">
          <v-card>
            <v-card-media :src="item.url" height="300px">
              <v-container fill-height fluid>
                <v-layout fill-height>
                  <v-flex xs12 align-end flexbox>
                    <span class="headline white--text" v-text="item.name"></span>
                  </v-flex>
                </v-layout>
              </v-container>
            </v-card-media>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn icon>
                <v-icon>delete</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <v-dialog v-model="dialog" max-width="350px">
      <v-btn fixed fab large bottom right class="mx-2 my-5" color="primary" dark slot="activator"><v-icon>add</v-icon></v-btn>
      <v-card>
        <v-card-title>
          <span class="headline">New face</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-layout column wrap align-center>
              <v-text-field label="Name" v-model="editedItem.name"></v-text-field>
              <file-input accept="image/*" ref="fileInput" @input="getUploadedFile" ></file-input>
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
</template>

<script>
  import FileInput from './FileInput.vue'

  export default {
    components: {FileInput},

    data: () => ({
      dialog: false,
      items: [],
      editedIndex: -1,
      editedItem: {
        name: '',
        url: ''
      },
      defaultItem: {
        name: '',
        url: ''
      }
    }),

    created () {
      this.initialize()
    },

    methods: {
      initialize () {
        this.$http.get(this.$store.state.endpoints.faces).then(response => {
          this.items = response.data
        }).catch(error => {
          console.log(error)
        })
      },

      getUploadedFile (e) {
        this.editedItem.url = e
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
        this.$http.post(this.$store.state.endpoints.faces, this.editedItem).then((response) => {
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
