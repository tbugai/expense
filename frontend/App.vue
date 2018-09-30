<template>
  <v-app>
  <v-navigation-drawer temporary></v-navigation-drawer>
  <v-toolbar app>
    <v-toolbar-title>Expense</v-toolbar-title>
  </v-toolbar>
  <v-content>
    <v-container fluid>
      <router-view></router-view>
    </v-container>
  </v-content>
  <v-footer app class="pa-3">
    <v-spacer></v-spacer>
    <div>&copy; {{ new Date().getFullYear() }}</div>
  </v-footer>
  </v-app>
</template>

<script>
export default {
  data () {
    return {
    }
  }
}
</script>

<style lang="scss">
</style>
