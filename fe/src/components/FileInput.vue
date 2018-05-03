<template>
  <div align="center">
    <div>
      <img
        :src="imageUrl"
        ref="imageUrl"
        height="150px"
        @click="onPickFile"
        style="cursor: pointer;"
        v-if="imageUrl"
      >
      <v-icon size="150px" v-else>face</v-icon>
    </div>
    <div>
      <v-btn @click="onPickFile" v-if="!imageUrl">
        {{ selectLabel }}
      </v-btn>
      <v-btn class="error" @click="removeFile" v-else>
        {{ removeLabel }}
      </v-btn>
      <input
        type="file"
        ref="image"
        name="image"
        :accept="accept"
        @change="onFilePicked"
      >
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      value: {
        type: String
      },
      accept: {
        type: String,
        default: '*'
      },
      selectLabel: {
        type: String,
        default: 'Select an image'
      },
      removeLabel: {
        type: String,
        default: 'Remove'
      }
    },

    data () {
      return {
        imageUrl: ''
      }
    },

    watch: {
      value (v) {
        this.imageUrl = v
      }
    },

    mounted () {
      this.imageUrl = this.value
    },

    methods: {
      onPickFile () {
        this.$refs.image.click()
      },

      onFilePicked (event) {
        const files = event.target.files || event.dataTransfer.files

        if (files && files[0]) {
          let filename = files[0].name

          if (filename && filename.lastIndexOf('.') <= 0) {
            return alert('Please add a valid image!')
          }

          const fileReader = new FileReader()
          fileReader.addEventListener('load', () => {
            this.imageUrl = fileReader.result
            this.$emit('input', this.imageUrl)
          })
          fileReader.readAsDataURL(files[0])
        }
      },

      removeFile () {
        this.imageUrl = ''
      }
    }
  }
</script>

<style scoped>
  input[type=file] {
    position: absolute;
    left: -99999px;
  }
</style>
