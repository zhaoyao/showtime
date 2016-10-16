<template>
  <div>
    <ul v-if="item.Episodes != null">
      <li v-for="epInfo in item.Episodes">
        {{epInfo.Name}}[{{epInfo.Format}}]

        <ul>
          <li v-for="link in epInfo.Links">
            <a :href="link.URL">{{link.Type}}</a>
          </li>
        </ul>
      </li>
    </ul>
    <div v-if="item.Episodes == null">No Links Available</div>
  </div>
</template>

<script>
export default {
  name: 'resource',
  props: ['item'],
  created () {
    console.log('Rendering item: ', this.item)
  }
}
</script>