<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="row">
      <q-toolbar
        class="col-8"
        style="background-color: rgb(16, 16, 16); color: gray"
      >
        <q-toolbar-title>Portfolio Manager</q-toolbar-title>
      </q-toolbar>
      <q-toolbar
        class="col-4"
        style="background-color: rgb(16, 16, 16); color: gray"
      >
        <q-toolbar-title style="text-align: right">{{
          dateString
        }}</q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'MainLayout',
  setup() {
    let currentTime = new Date();

    let dateString = '';

    dateString =
      currentTime.getMonth() +
      '/' +
      currentTime.getDay() +
      '/' +
      currentTime.getFullYear();

    return {
      dateString,
    };
  },
});
</script>
