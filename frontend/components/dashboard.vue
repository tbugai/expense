<template>
  <div>
    <v-layout justify-end>
      <v-btn color="primary" @click.native="dialog = true">Add Expense</v-btn>
    </v-layout>
    <expense-table></expense-table>
    <v-dialog v-model="dialog" max-width="290">
      <v-card>
        <v-card-title class="headline">Expense</v-card-title>
        <v-cart-text>
          <div class="pa-3">
          <v-form v-model="expense">
            <v-text-field
              v-model="description"
              label="Description"
              required></v-text-field>
            <v-text-field
              v-model="amount"
              label="Amount"
              type="number"
              prefix="$"
              required></v-text-field>
          </v-form>
          </div>
        </v-cart-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="secondary darken-1" flat="flat" @click.native="dialog = false">Cancel</v-btn>
          <v-btn color="primary darken-1" @click.native="dialog = false">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import ExpenseTableComponent from './expense-table.vue'

export default {
  data() {
    return {
      expense: {
        description: "",
        amount: ""
      },
      dialog: false
    }
  },

  components: {
    expenseTable: ExpenseTableComponent
  }
}
</script>
