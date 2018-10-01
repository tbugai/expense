<template>
  <expense-table></expense-table>
</template>

<script>
import ExpenseTableComponent from './expense-table.vue'

export default {
  data() {
    return {}
  },

  components: {
    expenseTable: ExpenseTableComponent
  }
}
</script>
