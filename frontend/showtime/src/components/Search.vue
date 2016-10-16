<template>
  <div>
    <form v-on:submit.prevent="search">
      <input name="q" v-model="q" >
    </form>
  </div>
</template>

<script>
export default {
  name: 'search',

  data () {
    return {
      q: ''
    }
  },

  beforeMount () {
    console.log('Search Mounted')
  },

  methods: {
    search () {
      console.log('Start searching: ', this.q)
      this.$router.push({
        'name': 'search',
        'query': { q: this.q }
      })
    }
  }
}
</script>
