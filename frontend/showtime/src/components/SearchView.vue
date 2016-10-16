<template>
  <div>
    <spinner :show="loading"></spinner>
    <resource-list v-if="!loading" :items="items"></resource-list>
  </div>
</template>

<script>
import api from '../api'
import ResourceList from './ResourceList'
import Spinner from './Spinner'

export default {
  name: 'SearchView',

  components: {ResourceList, Spinner},

  data () {
    return {
      loading: true,
      items: []
    }
  },

  created () {
    this.search()
  },

  watch: {
    '$route': 'search'
  },

  methods: {
    search () {
      console.log('SearchView: Search: ', this.$route.query)
      api.search(this.$route.query['q']).then((ret) => {
        console.log('search ret:', ret)
        this.items = ret.data
        this.loading = false
      })
    }
  }
}
</script>