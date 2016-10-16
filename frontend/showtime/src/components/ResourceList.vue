<template>
 <ul>
    <li v-for="item in items" :key="item.itemid">
      <router-link :to="{ name: 'resource', params: { id: item.itemid }}">
        <img :src="item.poster">
      </router-link>
      <router-link :to="{ name: 'resource', params: { id: item.itemid }}">
          {{item.title}} ({{item.itemid}})
      </router-link>
    </li>
  </ul>
</template>

<script>
import Resource from './Resource'
export default {
  name: 'resource-list',
  props: ['items'],
  components: [Resource]
}
</script>