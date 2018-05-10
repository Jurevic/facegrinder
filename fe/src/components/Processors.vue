<template>
  <div>
    <v-card>
      <v-card-text>
        <h1>Processors</h1>
      </v-card-text>
    </v-card>
    <v-flex class="text-xs-center" mt-5>
      <v-btn color="primary" @click="runProcessor">Run processor</v-btn>
    </v-flex>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    methods: {
      runProcessor () {
        axios.get(this.$store.state.endpoints.processors, {
          headers: {
            Authorization: this.$store.state.jwt
          }
        }).catch(error => {
          console.log(error)
        })
      }
    }
  }
</script>
