<template>
  <v-data-table
    :headers="headers"
    :items="expenses">
    <template slot="items" slot-scope="props">
        <td>{{ props.item.description }}</td>
        <td>${{ props.item.amount }}</td>
    </template>
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        { text: 'Description', value: 'description' },
        { text: 'Amount', value: 'amount' }
      ],
    }
  },

  mounted() {
    this.$store.dispatch("fetch-expenses")
  },

  computed: {
    expenses() {
      return this.$store.getters.getExpenses
    }
  }
}
</script>
