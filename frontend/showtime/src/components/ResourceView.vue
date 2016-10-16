<template>
  <div>
    <spinner :show="loading"></spinner>
    <resource v-if="item != null" :item="item"></resource>
  </div>
</template>


<script>
import Resource from './Resource'
import Spinner from './Spinner'
import api from '../api'

export default {
  components: {Spinner, Resource},
  data () {
    return {
      loading: true,
      item: null
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      let id = this.$route.params['id']
      api.getResource(id)
        .then((ret) => {
          this.loading = false
          this.item = ret.data
        })
    }
  }
}
</script>