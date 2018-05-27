<template>
  <div>
    <v-navigation-drawer right absolute temporary v-model="drawer">
      <v-list class="pa-1">
        <v-list-tile>
          <v-list-tile-content>
            <v-list-tile-title><h2>Select processor</h2></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
      <v-list class="pa-1">
        <v-divider></v-divider>
        <v-list-tile v-for="(item, index) in items" :key="item.name" @click="selected(index)">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.name }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <div align="center">
          <v-dialog v-model="dialog" max-width="500px">
            <v-btn color="primary" dark slot="activator">Add processor</v-btn>
            <v-card>
              <v-card-title>
                <span class="headline">New processor</span>
              </v-card-title>
              <v-card-text>
                <v-container grid-list-md>
                  <v-layout wrap>
                    <v-flex>
                      <v-text-field label="Processor name" v-model="editedItem.name"></v-text-field>
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
        <h1>Processors</h1>
      </v-card-text>
    </v-card>
    <v-container grid-list-md>
      <v-layout row>
        <v-flex>
        </v-flex>
        <v-flex xs6>
          <v-flex class="text-xs-center" pa-3>
            <h1>{{selectedItem.name}}</h1>
          </v-flex>
          <v-layout column v-for="(item, index) in selectedItem.nodes" :key="item.key">
            <v-flex>
              <v-card>
                <v-card-title>
                  <span class="headline">{{getElemName(item.key)}}</span>
                </v-card-title>
                <v-card-media :src="item.url">
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
                    <v-icon>edit</v-icon>
                  </v-btn>
                  <v-btn icon>
                    <v-icon>delete</v-icon>
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-flex>
            <v-flex class="text-xs-center">
              <div align="center">
                <v-dialog v-model="elemDialog" max-width="500px">
                  <v-btn icon large slot="activator" v-if="index !== selectedItem.nodes.length - 1"><v-icon>arrow_downward</v-icon></v-btn>
                  <v-btn icon fab large slot="activator" v-if="index === selectedItem.nodes.length - 1"><v-icon>add</v-icon></v-btn>
                  <v-card>
                    <v-card-title>
                      <span class="headline">Add element</span>
                    </v-card-title>
                    <v-card-text>
                      <v-container grid-list-md>
                        <v-layout wrap>
                          <v-flex xs12>
                            <v-select
                              :items="elemChoices"
                              v-model="selectedElemKey"
                              label="Select element type"
                              class="input-group--focused"
                            ></v-select>
                          </v-flex>
                          <div v-if="selectedElemKey" width=100%>
                            <v-flex xs12 v-for="(value, key) in elemTypes" :key="key">
                              <v-text-field
                                :name="key"
                                :label="key"
                                :type="value"
                                :value="getElemDefault(key)"
                              ></v-text-field>
                            </v-flex>
                          </div>
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
            </v-flex>
          </v-layout>
          <v-flex class="text-xs-center">
            <v-btn color="primary" @click="runProcessor">Run processor</v-btn>
          </v-flex>
        </v-flex>
        <v-flex xs3>
        </v-flex>
      </v-layout>
    </v-container>
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
  export default {
    data: () => ({
      drawer: false,
      keyVisible: false,
      dialog: false,
      elemDialog: false,
      items: [],
      choices: [],
      editedIndex: -1,
      editedElem: {},
      selectedElemKey: '',
      editedItem: {
        id: null,
        name: '',
        nodes: []
      },
      selectedItem: {
        id: null,
        name: '',
        nodes: []
      }
    }),

    created () {
      this.initialize()
    },

    computed: {
      elemChoices () {
        let choices = []

        for (let key in this.choices) {
          choices.push({text: this.choices[key].name, value: key})
        }

        return choices
      },

      elemTypes () {
        return this.choices[this.selectedElemKey].types
      }
    },

    methods: {
      initialize () {
        this.getProcessors()
        this.getProcessorChoices()
      },

      getElemName (key) {
        let name = this.choices[key].name
        if (!name) {
          name = 'Unknown'
        }

        return name
      },

      getElemDefault (key) {
        return this.choices[this.selectedElemKey].params[key]
      },

      addElem (index) {
        console.log('add')
      },

      close () {
        this.dialog = false
        this.elemDialog = false
        setTimeout(() => {
          this.editedItem = {}
          this.editedElem = {}
          this.editedIndex = -1
        }, 300)
      },

      save () {
        this.dialog = false
      },

      selected (index) {
        this.selectedItem = this.items[index]
      },

      runProcessor () {
        this.$http.get(this.$store.state.endpoints.processors + this.selectedItem.id + '/run').catch(error => {
          console.log(error)
        })
      },

      getProcessors () {
        this.$http.get(this.$store.state.endpoints.processors).then(response => {
          this.items = response.data
          this.selectedItem = this.items[0]
        }).catch(error => {
          console.log(error)
        })
      },

      getProcessorChoices () {
        this.choices = this.$store.state.processorChoices

        if (this.choices.length <= 0) {
          this.$http.get(this.$store.state.endpoints.processorChoices).then(response => {
            this.$store.commit('setProcessorChoices', response.data)
            this.choices = response.data
          }).catch(error => {
            console.log(error)
          })
        }
      }
    }
  }
</script>
